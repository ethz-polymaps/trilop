package trilateration

import (
	"errors"
	"math"

	"gonum.org/v1/gonum/optimize"

	"github.com/ethz-polymaps/trilop"
	"github.com/ethz-polymaps/trilop/distance"
)

type (
	Measurements []Measurement
	Measurement  struct {
		Lat      float64
		Lon      float64
		Distance float64
		Weight   float64
	}
)

func Trilaterate(measurements []Measurement) (loc trilop.Position, accuracy float64, err error) {

	if len(measurements) < 1 || len(measurements) > 3 {
		return trilop.EmptyPosition, 0, errors.New("must provide 1-3 measurements")
	}

	if len(measurements) == 1 {
		return trilop.NewPosition(measurements[0].Lat, measurements[0].Lon), measurements[0].Distance, nil
	}

	for _, m := range measurements {
		if m.Weight <= 0 {
			return trilop.EmptyPosition, 0, errors.New("weights must be positive")
		}
	}

	for _, m := range measurements {
		if m.Distance < 0 {
			return trilop.EmptyPosition, 0, errors.New("distances must be positive")
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
				// Calculate distance between current point and measurement
				d := distance.Calculate(trilop.NewPosition(lat, lon), trilop.NewPosition(measurement.Lat, measurement.Lon))
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
		return trilop.EmptyPosition, 0, err
	}

	weightedSquareError := result.F
	weightedError := math.Sqrt(weightedSquareError) / float64(len(measurements))
	return trilop.NewPosition(result.X[0], result.X[1]), weightedError, nil
}
