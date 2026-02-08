package polaris_test

import (
	"fmt"

	"github.com/ethz-polymaps/polaris"
)

func ExampleNewPosition() {
	// Create a position for Zurich, Switzerland
	zurich := polaris.NewPosition(47.3769, 8.5417)
	fmt.Printf("Latitude: %.4f\n", zurich.Latitude)
	fmt.Printf("Longitude: %.4f\n", zurich.Longitude)
	// Output:
	// Latitude: 47.3769
	// Longitude: 8.5417
}

func ExamplePosition_String() {
	pos := polaris.NewPosition(47.3769, 8.5417)
	fmt.Println(pos)
	// Output:
	// 47.376900,8.541700
}
