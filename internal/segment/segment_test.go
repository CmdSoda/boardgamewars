package segment

import (
	"github.com/CmdSoda/boardgamewars/internal/math"
	"reflect"
	"testing"
)

func TestNewSegment(t *testing.T) {
	type args struct {
		direction Direction
		center    math.Vector2
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
				center:    math.Vector2{},
			},
			want: Segment{
				Start: math.Vector2{
					X: 0,
					Y: -1,
				},
				End: math.Vector2{
					X: math.Hr,
					Y: -0.5,
				},
			},
		},
		{
			name: "E",
			args: args{
				direction: E,
				center:    math.Vector2{},
			},
			want: Segment{
				Start: math.Vector2{
					X: math.Hr,
					Y: -0.5,
				},
				End: math.Vector2{
					X: math.Hr,
					Y: 0.5,
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
