package hexagon

import (
	"github.com/CmdSoda/boardgamewars/internal/vector"
	"reflect"
	"testing"
)

func TestNewSegment(t *testing.T) {
	type args struct {
		direction Direction
		center    vector.Vector
	}
	tests := []struct {
		name string
		args args
		want Segment
	}{
		{
			name: "NE",
			args: args{
				direction: NE,
				center:    vector.Vector{},
			},
			want: Segment{
				Start: vector.Vector{
					X: 0,
					Y: -1,
				},
				End: vector.Vector{
					X: HexHalfWidth,
					Y: -0.5,
				},
			},
		},
		{
			name: "E",
			args: args{
				direction: E,
				center:    vector.Vector{},
			},
			want: Segment{
				Start: vector.Vector{
					X: HexHalfWidth,
					Y: -0.5,
				},
				End: vector.Vector{
					X: HexHalfWidth,
					Y: 0.5,
				},
			},
		},
		{
			name: "SE",
			args: args{
				direction: SE,
				center:    vector.Vector{},
			},
			want: Segment{
				Start: vector.Vector{
					X: HexHalfWidth,
					Y: 0.5,
				},
				End: vector.Vector{
					X: 0,
					Y: 1,
				},
			},
		},
		{
			name: "SW",
			args: args{
				direction: SW,
				center:    vector.Vector{},
			},
			want: Segment{
				Start: vector.Vector{
					X: 0,
					Y: 1,
				},
				End: vector.Vector{
					X: -HexHalfWidth,
					Y: 0.5,
				},
			},
		},
		{
			name: "W",
			args: args{
				direction: W,
				center:    vector.Vector{},
			},
			want: Segment{
				Start: vector.Vector{
					X: -HexHalfWidth,
					Y: 0.5,
				},
				End: vector.Vector{
					X: -HexHalfWidth,
					Y: -0.5,
				},
			},
		},
		{
			name: "NW",
			args: args{
				direction: NW,
				center:    vector.Vector{},
			},
			want: Segment{
				Start: vector.Vector{
					X: -HexHalfWidth,
					Y: -0.5,
				},
				End: vector.Vector{
					X: 0,
					Y: -1,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSegment(tt.args.direction, tt.args.center); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSegment() = %v, want %v", got, tt.want)
			}
		})
	}
}
