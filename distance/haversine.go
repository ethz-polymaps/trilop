package distance

import (
	"math"

	"github.com/ethz-polymaps/polaris"
)

// HaversineDistance calculates the great-circle distance between two points
// on Earth using the Haversine formula. This assumes a spherical Earth with
// a radius of 6,371 km. Returns the distance in meters.
//
// The Haversine formula is computationally efficient and provides good accuracy
// for most applications. For higher precision over long distances, consider
// using [VincentyDistance] instead.
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
