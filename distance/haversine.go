package distance

import (
	"math"

	"github.com/ethz-polymaps/polaris"
)

// HaversineDistance calculates the distance between two points on Earth
func HaversineDistance(a, b polaris.Position) float64 {
	const R = 6371000 // Earth radius in meters

	lat1Rad := a.Latitude * math.Pi / 180
	lat2Rad := b.Latitude * math.Pi / 180
	deltaLat := (b.Latitude - a.Latitude) * math.Pi / 180
	deltaLon := (b.Longitude - a.Longitude) * math.Pi / 180

	x := math.Sin(deltaLat/2)*math.Sin(deltaLat/2) +
		math.Cos(lat1Rad)*math.Cos(lat2Rad)*
			math.Sin(deltaLon/2)*math.Sin(deltaLon/2)

	return R * 2 * math.Atan2(math.Sqrt(x), math.Sqrt(1-x))
}
