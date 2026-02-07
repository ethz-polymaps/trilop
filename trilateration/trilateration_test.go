package trilateration

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTrilaterate(t *testing.T) {

	measurements := []Measurement{
		{Lat: 47.41331043239206, Lon: 8.536443579900189, Distance: 0.7686246100397739, Weight: 1.0},
		{Lat: 47.41321841412086, Lon: 8.536437101250389, Distance: 0.8767123872968682, Weight: 1.0},
		{Lat: 47.41330944456364, Lon: 8.536520373280595, Distance: 1.4839817889675653, Weight: 1.0},
	}

	loc, accuracy, err := Trilaterate(measurements)
	require.NoError(t, err)

	assert.InDelta(t, 8.536464538143443, loc.Longitude, 0.0000000000009)
	assert.InDelta(t, 47.413276910646324, loc.Latitude, 0.0000000000009)
	assert.InDelta(t, 2.637536791157061, accuracy, 0.0000000000009)
}
