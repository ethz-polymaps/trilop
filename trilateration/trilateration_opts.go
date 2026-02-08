package trilateration

import "github.com/ethz-polymaps/polaris"

// WithDistanceFunc sets the distance calculation function used by the Trilaterator.
// Use this to switch from the default Haversine formula to Vincenty for higher accuracy:
//
//	t := NewTrilaterator(WithDistanceFunc(distance.VincentyDistance))
func WithDistanceFunc(distanceFunc func(a polaris.Position, b polaris.Position) float64) TrilateratorOpt {
	return func(t *TrilateratorConfig) {
		t.DistanceFunc = distanceFunc
	}
}
