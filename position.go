package polaris

import "fmt"

type (
	Position struct {
		Latitude  float64
		Longitude float64
	}
)

var (
	EmptyPosition Position
)

func NewPosition(latitude float64, longitude float64) Position {
	return Position{Latitude: latitude, Longitude: longitude}
}

func (l Position) String() string {
	return fmt.Sprintf("%f,%f", l.Latitude, l.Longitude)
}
