// Package trilateration provides position estimation from multiple distance measurements.
//
// Trilateration is a technique for determining a location by measuring distances
// from known reference points. Unlike triangulation (which uses angles), trilateration
// uses distances to compute the intersection point.
//
// # How It Works
//
// Given a set of reference points (e.g., beacons, cell towers, WiFi access points)
// with known positions and measured distances to an unknown target, the algorithm
// finds the position that best fits all measurements.
//
// The implementation uses the Nelder-Mead optimization algorithm to minimize
// the weighted sum of squared distance errors. This approach is robust to
// measurement noise and can handle over-determined systems (more than 3 measurements).
//
// # Usage
//
// Create a [Trilaterator] and call [Trilaterator.Trilaterate] with measurements:
//
//	t := trilateration.NewTrilaterator()
//	measurements := []trilateration.Measurement{
//	    {Lat: 47.4133, Lon: 8.5364, Distance: 1000, Weight: 1.0},
//	    {Lat: 47.4200, Lon: 8.5400, Distance: 800, Weight: 1.0},
//	    {Lat: 47.4100, Lon: 8.5500, Distance: 1200, Weight: 1.0},
//	}
//	position, accuracy, err := t.Trilaterate(measurements)
//
// # Weights
//
// Each measurement includes a Weight field that indicates the confidence level.
// Higher weights give more influence to that measurement. Use lower weights for:
//   - Measurements with known higher noise
//   - Less reliable signal sources
//   - Older measurements in time-series data
//
// # Configuration
//
// The default [Trilaterator] uses [distance.HaversineDistance] for distance calculations.
// For higher accuracy, use [WithDistanceFunc] to specify [distance.VincentyDistance]:
//
//	t := trilateration.NewTrilaterator(
//	    trilateration.WithDistanceFunc(distance.VincentyDistance),
//	)
package trilateration
