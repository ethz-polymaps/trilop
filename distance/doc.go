// Package distance provides functions for calculating distances between geographic positions.
//
// This package implements two distance calculation methods:
//
// # Haversine Formula
//
// [HaversineDistance] calculates the great-circle distance between two points
// on a sphere. It assumes a perfectly spherical Earth with radius 6,371 km.
// This method is computationally efficient and suitable for most applications
// where high precision is not critical.
//
// Accuracy: ~0.5% error compared to the true geodesic distance.
//
// # Vincenty Formula
//
// [VincentyDistance] calculates the geodesic distance using Vincenty's formulae,
// which model the Earth as an oblate spheroid (WGS-84 ellipsoid). This method
// is more accurate than Haversine, especially for long distances and paths
// that cross near the poles.
//
// Accuracy: Sub-millimeter precision for any distance on Earth.
//
// # Choosing a Method
//
// Use [HaversineDistance] when:
//   - Performance is critical
//   - Distances are relatively short (< 100 km)
//   - Approximate results are acceptable
//
// Use [VincentyDistance] when:
//   - High accuracy is required
//   - Calculating long distances
//   - Working with surveying or mapping applications
//
// Both functions return the distance in meters.
package distance
