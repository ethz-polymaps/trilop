# Polaris

[![Test](https://github.com/ethz-polymaps/polaris/actions/workflows/test.yml/badge.svg)](https://github.com/ethz-polymaps/polaris/actions/workflows/test.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/ethz-polymaps/polaris.svg)](https://pkg.go.dev/github.com/ethz-polymaps/polaris)

Polaris is a Go library for geolocation calculations, providing accurate distance measurements and trilateration for position estimation.

## Features

- **Distance Calculations**: Compute distances between geographic coordinates using:
  - **Haversine formula**: Fast calculation assuming a spherical Earth
  - **Vincenty formula**: High-precision calculation using the WGS-84 ellipsoid model
- **Trilateration**: Estimate position from multiple distance measurements using weighted least-squares optimization

## Installation

```bash
go get github.com/ethz-polymaps/polaris
```

## Quick Start

```go
package main

import (
    "fmt"

    "github.com/ethz-polymaps/polaris"
    "github.com/ethz-polymaps/polaris/distance"
)

func main() {
    zurich := polaris.NewPosition(47.3769, 8.5417)
    bern := polaris.NewPosition(46.9480, 7.4474)

    dist := distance.HaversineDistance(zurich, bern)
    fmt.Printf("Distance: %.2f km\n", dist/1000)
    // Output: Distance: 95.49 km
}
```

See the [Go documentation](https://pkg.go.dev/github.com/ethz-polymaps/polaris) for complete API reference and runnable examples.

## Packages

### `polaris`

Core types for geographic positions.

```go
pos := polaris.NewPosition(47.3769, 8.5417)
fmt.Println(pos.Latitude, pos.Longitude)
```

### `polaris/distance`

Distance calculation functions. Both return distance in **meters**.

| Function | Model | Use Case |
|----------|-------|----------|
| `HaversineDistance` | Spherical Earth | Fast calculations, ~0.5% error |
| `VincentyDistance` | WGS-84 Ellipsoid | High precision, sub-millimeter accuracy |

### `polaris/trilateration`

Position estimation from distance measurements using the Nelder-Mead optimization algorithm.

```go
t := trilateration.NewTrilaterator()
measurements := []trilateration.Measurement{
    {Lat: 47.4133, Lon: 8.5364, Distance: 500, Weight: 1.0},
    {Lat: 47.4100, Lon: 8.5400, Distance: 300, Weight: 1.0},
    {Lat: 47.4120, Lon: 8.5450, Distance: 400, Weight: 1.0},
}
position, accuracy, err := t.Trilaterate(measurements)
```

Use `WithDistanceFunc(distance.VincentyDistance)` for higher accuracy.

## Contributing

Contributions are welcome! Please feel free to submit issues and pull requests.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
