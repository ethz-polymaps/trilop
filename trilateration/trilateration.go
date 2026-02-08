package trilateration

import (
	"errors"
	"fmt"
	"math"

	"github.com/ethz-polymaps/polaris/distance"
	"gonum.org/v1/gonum/optimize"

	"github.com/ethz-polymaps/polaris"
)

// Measurements is a slice of Measurement values.
type Measurements []Measurement

// Measurement represents a distance observation from a known reference point.
// It contains the reference point's coordinates, the measured distance to the
// unknown target position, and a weight indicating the measurement's reliability.
type Measurement struct {
	// Lat is the latitude of the reference point in decimal degrees.
	Lat float64
	// Lon is the longitude of the reference point in decimal degrees.
	Lon float64
	// Distance is the measured distance from the reference point to the target in meters.
	Distance float64
	// Weight indicates the measurement's reliability (higher = more trusted).
	// Must be positive. Use lower weights for noisier or less reliable measurements.
	Weight float64
}

// DistanceFunc is a function that calculates the distance between two positions.
// The distance package provides [distance.HaversineDistance] and [distance.VincentyDistance].
type DistanceFunc func(a, b polaris.Position) float64

// TrilateratorOpt is a functional option for configuring a Trilaterator.
type TrilateratorOpt func(*TrilateratorConfig)

// Trilaterator estimates geographic positions using trilateration.
type Trilaterator struct {
	config *TrilateratorConfig
}

// TrilateratorConfig holds the configuration for a Trilaterator.
type TrilateratorConfig struct {
	// DistanceFunc is used to calculate distances during optimization.
	// Defaults to distance.HaversineDistance.
	DistanceFunc DistanceFunc
	// MinMeasurements is the maximum number of measurements allowed.
	// Defaults to 3.
	MinMeasurements int
}

// NewTrilaterator creates a new Trilaterator with the given options.
// By default, it uses [distance.HaversineDistance] for distance calculations
// and allows up to 3 measurements.
func NewTrilaterator(opts ...TrilateratorOpt) *Trilaterator {
	config := &TrilateratorConfig{
		DistanceFunc:    distance.HaversineDistance,
		MinMeasurements: 3,
	}

	for _, opt := range opts {
		opt(config)
	}

	return &Trilaterator{
		config: config,
	}
}

// Trilaterate estimates a position from the given distance measurements using
// weighted least-squares optimization. It returns the estimated position,
// an accuracy metric (weighted RMS error in meters), and any error encountered.
//
// The function requires 1 to MinMeasurements measurements. With a single measurement,
// it returns the measurement's position with its distance as the accuracy.
// With multiple measurements, it uses the Nelder-Mead algorithm to find the
// position that minimizes the weighted sum of squared distance errors.
//
// All measurements must have positive weights and non-negative distances.
func (t *Trilaterator) Trilaterate(measurements []Measurement) (loc polaris.Position, accuracy float64, err error) {

	if len(measurements) < 1 || len(measurements) > t.config.MinMeasurements {
		return polaris.EmptyPosition, 0, fmt.Errorf("must provide 1-%d measurements", t.config.MinMeasurements)
	}

	if len(measurements) == 1 {
		return polaris.NewPosition(measurements[0].Lat, measurements[0].Lon), measurements[0].Distance, nil
	}

	for _, m := range measurements {
		if m.Weight <= 0 {
			return polaris.EmptyPosition, 0, errors.New("weights must be positive")
		}
	}

	for _, m := range measurements {
		if m.Distance < 0 {
			return polaris.EmptyPosition, 0, errors.New("distances must be positive")
		}
	}

	// Initial guess: average of beacon positions
	initLat, initLong := 0.0, 0.0
	totalWeight := 0.0
	for _, b := range measurements {
		initLat += b.Lat * b.Weight
		initLong += b.Lon * b.Weight
		totalWeight += b.Weight
	}
	initLat /= totalWeight
	initLong /= totalWeight

	problem := optimize.Problem{
		Func: func(x []float64) float64 {
			lat, lon := x[0], x[1]
			var sum float64
			for _, measurement := range measurements {
				// HaversineDistance distance between current point and measurement
				d := t.config.DistanceFunc(polaris.NewPosition(lat, lon), polaris.NewPosition(measurement.Lat, measurement.Lon))
				// Weighted square error
				diff := d - measurement.Distance
				sum += measurement.Weight * diff * diff
			}
			return sum
		},
	}

	// Run optimization
	result, err := optimize.Minimize(problem, []float64{initLat, initLong}, nil, &optimize.NelderMead{})
	if err != nil {
		return polaris.EmptyPosition, 0, err
	}

	weightedSquareError := result.F
	weightedError := math.Sqrt(weightedSquareError) / float64(len(measurements))
	return polaris.NewPosition(result.X[0], result.X[1]), weightedError, nil
}
