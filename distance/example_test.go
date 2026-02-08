package distance_test

import (
	"fmt"

	"github.com/ethz-polymaps/polaris"
	"github.com/ethz-polymaps/polaris/distance"
)

func ExampleHaversineDistance() {
	zurich := polaris.NewPosition(47.3769, 8.5417)
	bern := polaris.NewPosition(46.9480, 7.4474)

	dist := distance.HaversineDistance(zurich, bern)
	fmt.Printf("Distance: %.2f km\n", dist/1000)
	// Output:
	// Distance: 95.49 km
}

func ExampleVincentyDistance() {
	zurich := polaris.NewPosition(47.3769, 8.5417)
	bern := polaris.NewPosition(46.9480, 7.4474)

	dist := distance.VincentyDistance(zurich, bern)
	fmt.Printf("Distance: %.2f km\n", dist/1000)
	// Output:
	// Distance: 95.70 km
}

func Example_compareDistanceMethods() {
	// Comparing Haversine vs Vincenty over a long distance
	newYork := polaris.NewPosition(40.7128, -74.0060)
	london := polaris.NewPosition(51.5074, -0.1278)

	haversine := distance.HaversineDistance(newYork, london)
	vincenty := distance.VincentyDistance(newYork, london)

	fmt.Printf("Haversine: %.2f km\n", haversine/1000)
	fmt.Printf("Vincenty:  %.2f km\n", vincenty/1000)
	fmt.Printf("Difference: %.2f km\n", (haversine-vincenty)/1000)
	// Output:
	// Haversine: 5570.22 km
	// Vincenty:  5585.23 km
	// Difference: -15.01 km
}
