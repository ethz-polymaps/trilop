// Package polaris provides geolocation primitives and calculations for Go.
//
// Polaris offers accurate distance measurements between geographic coordinates
// and trilateration capabilities for position estimation from multiple reference points.
//
// The core type is [Position], which represents a geographic location using
// latitude and longitude coordinates in decimal degrees.
//
// # Distance Calculations
//
// The distance subpackage provides two methods for calculating distances:
//
//   - [distance.HaversineDistance]: Uses the Haversine formula assuming a spherical Earth.
//     Faster but less accurate for long distances.
//
//   - [distance.VincentyDistance]: Uses Vincenty's formulae with the WGS-84 ellipsoid model.
//     More accurate, especially for long distances.
//
// # Trilateration
//
// The trilateration subpackage estimates a position from multiple distance measurements
// using weighted least-squares optimization. This is useful for applications like:
//
//   - Indoor positioning systems
//   - GPS-denied navigation
//   - Beacon-based localization
//
// # Example
//
//	zurich := polaris.NewPosition(47.3769, 8.5417)
//	bern := polaris.NewPosition(46.9480, 7.4474)
//	dist := distance.HaversineDistance(zurich, bern)
//	fmt.Printf("Distance: %.2f km\n", dist/1000)
package polaris
