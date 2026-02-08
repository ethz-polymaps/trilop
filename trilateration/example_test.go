package trilateration_test

import (
	"fmt"

	"github.com/ethz-polymaps/polaris/distance"
	"github.com/ethz-polymaps/polaris/trilateration"
)

func ExampleTrilaterator_Trilaterate() {
	// Create a trilaterator with default settings (uses Haversine distance)
	t := trilateration.NewTrilaterator()

	// Define measurements from known reference points
	// Each measurement includes: latitude, longitude, distance to target (meters), weight
	measurements := []trilateration.Measurement{
		{Lat: 47.4133, Lon: 8.5364, Distance: 500, Weight: 1.0},
		{Lat: 47.4100, Lon: 8.5400, Distance: 300, Weight: 1.0},
		{Lat: 47.4120, Lon: 8.5450, Distance: 400, Weight: 1.0},
	}

	position, accuracy, err := t.Trilaterate(measurements)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Position: %.4f, %.4f\n", position.Latitude, position.Longitude)
	fmt.Printf("Accuracy: %.2f meters\n", accuracy)
	// Output:
	// Position: 47.4094, 8.5423
	// Accuracy: 57.54 meters
}

func ExampleNewTrilaterator_withVincenty() {
	// Create a trilaterator with Vincenty distance for higher accuracy
	t := trilateration.NewTrilaterator(
		trilateration.WithDistanceFunc(distance.VincentyDistance),
	)

	measurements := []trilateration.Measurement{
		{Lat: 47.4133, Lon: 8.5364, Distance: 500, Weight: 1.0},
		{Lat: 47.4100, Lon: 8.5400, Distance: 300, Weight: 1.0},
		{Lat: 47.4120, Lon: 8.5450, Distance: 400, Weight: 1.0},
	}

	position, accuracy, err := t.Trilaterate(measurements)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Position: %.4f, %.4f\n", position.Latitude, position.Longitude)
	fmt.Printf("Accuracy: %.2f meters\n", accuracy)
	// Output:
	// Position: 47.4095, 8.5423
	// Accuracy: 57.60 meters
}

func ExampleTrilaterator_Trilaterate_withWeights() {
	t := trilateration.NewTrilaterator()

	// Use weights to indicate measurement confidence
	// Higher weight = more influence on the result
	measurements := []trilateration.Measurement{
		{Lat: 47.4133, Lon: 8.5364, Distance: 500, Weight: 1.0}, // Normal confidence
		{Lat: 47.4100, Lon: 8.5400, Distance: 300, Weight: 2.0}, // High confidence
		{Lat: 47.4120, Lon: 8.5450, Distance: 400, Weight: 0.5}, // Low confidence
	}

	position, accuracy, err := t.Trilaterate(measurements)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Position: %.4f, %.4f\n", position.Latitude, position.Longitude)
	fmt.Printf("Accuracy: %.2f meters\n", accuracy)
	// Output:
	// Position: 47.4124, 8.5418
	// Accuracy: 45.52 meters
}
