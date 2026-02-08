package trilateration

import "github.com/ethz-polymaps/polaris"

func WithDistanceFunc(vincentyDistance func(a polaris.Position, b polaris.Position) float64) TrilateratorOpt {
	return func(t *TrilateratorConfig) {
		t.DistanceFunc = vincentyDistance
	}
}
