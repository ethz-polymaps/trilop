package polaris

import "fmt"

// Position represents a geographic location using latitude and longitude
// coordinates in decimal degrees. Latitude ranges from -90 (South) to +90 (North),
// and longitude ranges from -180 (West) to +180 (East).
type Position struct {
	Latitude  float64
	Longitude float64
}

// EmptyPosition is the zero value for Position, representing the coordinates (0, 0)
// at the intersection of the Prime Meridian and the Equator in the Gulf of Guinea.
var EmptyPosition Position

// NewPosition creates a new Position with the given latitude and longitude
// in decimal degrees.
func NewPosition(latitude float64, longitude float64) Position {
	return Position{Latitude: latitude, Longitude: longitude}
}

// String returns a string representation of the position in "latitude,longitude" format.
func (l Position) String() string {
	return fmt.Sprintf("%f,%f", l.Latitude, l.Longitude)
}
