package trilateration

import (
	"errors"
	"fmt"
	"math"

	"github.com/ethz-polymaps/polaris/distance"
	"gonum.org/v1/gonum/optimize"

	"github.com/ethz-polymaps/polaris"
)

type (
	Measurements []Measurement
	Measurement  struct {
		Lat      float64
		Lon      float64
		Distance float64
		Weight   float64
	}

	DistanceFunc    func(a, b polaris.Position) float64
	TrilateratorOpt func(*TrilateratorConfig)
	Trilaterator    struct {
		config *TrilateratorConfig
	}
	TrilateratorConfig struct {
		DistanceFunc    DistanceFunc
		MinMeasurements int
	}
)

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

// Trilaterate calculates the position and accuracy of a device based on trilateration
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
