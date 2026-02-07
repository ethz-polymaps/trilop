package distance

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ethz-polymaps/trilop"
)

func TestCalculateHaversine(t *testing.T) {
	type args struct {
		a trilop.Position
		b trilop.Position
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "~0m",
			args: args{
				a: trilop.NewPosition(47.413310, 8.536444),
				b: trilop.NewPosition(47.413310, 8.536444),
			},
			want: 0,
		},
		{
			name: "~5.7m",
			args: args{
				a: trilop.NewPosition(47.413310, 8.536444),
				b: trilop.NewPosition(47.413309, 8.536520),
			},
			want: 5.719788976313551,
		},
		{
			name: "~5.7m",
			args: args{
				a: trilop.NewPosition(47.463960, 8.322321),
				b: trilop.NewPosition(47.474113, 8.305055),
			},
			want: 1720.1466235357145,
		},
		{
			name: "~382900m",
			args: args{
				b: trilop.NewPosition(39.099912, -94.581213),
				a: trilop.NewPosition(38.627089, -90.200203),
			},
			want: 382900.0503756016,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Calculate(tt.args.a, tt.args.b)
			assert.InDelta(t, tt.want, result, 0.000000000000001)
		})
	}
}
