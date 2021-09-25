package hexagon

import (
	"github.com/CmdSoda/boardgamewars/internal/math"
	"reflect"
	"testing"
)

func TestGetCenterCoordinates(t *testing.T) {
	type args struct {
		h hexagon
	}
	tests := []struct {
		name string
		args args
		want math.Vector2
	}{
		{"0", args{h: hexagon{
			column: 0,
			row:    0,
		}}, math.Vector2{X: 0.8660254037844386, Y: 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetCenterCoordinates(tt.args.h); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCenterCoordinates() = %v, want %v", got, tt.want)
			}
		})
	}
}
