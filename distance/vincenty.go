package distance

import (
	"math"

	"github.com/ethz-polymaps/polaris"
)

// WGS-84 ellipsoid parameters
const (
	wgs84A = 6378137.0         // Semi-major axis in meters
	wgs84B = 6356752.314245    // Semi-minor axis in meters
	wgs84F = 1 / 298.257223563 // Flattening
)

// VincentyDistance calculates the distance between two points using Vincenty's formula.
// This is more accurate than Haversine as it accounts for Earth's ellipsoidal shape (WGS-84).
// Returns distance in meters.
func VincentyDistance(a, b polaris.Position) float64 {
	if a.Latitude == b.Latitude && a.Longitude == b.Longitude {
		return 0
	}

	lat1 := a.Latitude * math.Pi / 180
	lat2 := b.Latitude * math.Pi / 180
	lon1 := a.Longitude * math.Pi / 180
	lon2 := b.Longitude * math.Pi / 180

	L := lon2 - lon1
	U1 := math.Atan((1 - wgs84F) * math.Tan(lat1))
	U2 := math.Atan((1 - wgs84F) * math.Tan(lat2))

	sinU1, cosU1 := math.Sin(U1), math.Cos(U1)
	sinU2, cosU2 := math.Sin(U2), math.Cos(U2)

	lambda := L
	var sinSigma, cosSigma, sigma, sinAlpha, cos2Alpha, cos2SigmaM float64

	for i := 0; i < 100; i++ {
		sinLambda, cosLambda := math.Sin(lambda), math.Cos(lambda)

		sinSigma = math.Sqrt(
			(cosU2*sinLambda)*(cosU2*sinLambda) +
				(cosU1*sinU2-sinU1*cosU2*cosLambda)*(cosU1*sinU2-sinU1*cosU2*cosLambda))

		if sinSigma == 0 {
			return 0 // Co-incident points
		}

		cosSigma = sinU1*sinU2 + cosU1*cosU2*cosLambda
		sigma = math.Atan2(sinSigma, cosSigma)

		sinAlpha = cosU1 * cosU2 * sinLambda / sinSigma
		cos2Alpha = 1 - sinAlpha*sinAlpha
		cos2SigmaM = cosSigma - 2*sinU1*sinU2/cos2Alpha

		if math.IsNaN(cos2SigmaM) {
			cos2SigmaM = 0 // Equatorial line
		}

		C := wgs84F / 16 * cos2Alpha * (4 + wgs84F*(4-3*cos2Alpha))
		lambdaPrev := lambda
		lambda = L + (1-C)*wgs84F*sinAlpha*
			(sigma+C*sinSigma*(cos2SigmaM+C*cosSigma*(-1+2*cos2SigmaM*cos2SigmaM)))

		if math.Abs(lambda-lambdaPrev) < 1e-12 {
			break
		}
	}

	uSq := cos2Alpha * (wgs84A*wgs84A - wgs84B*wgs84B) / (wgs84B * wgs84B)
	A := 1 + uSq/16384*(4096+uSq*(-768+uSq*(320-175*uSq)))
	B := uSq / 1024 * (256 + uSq*(-128+uSq*(74-47*uSq)))

	deltaSigma := B * sinSigma * (cos2SigmaM + B/4*(cosSigma*(-1+2*cos2SigmaM*cos2SigmaM)-
		B/6*cos2SigmaM*(-3+4*sinSigma*sinSigma)*(-3+4*cos2SigmaM*cos2SigmaM)))

	return wgs84B * A * (sigma - deltaSigma)
}
