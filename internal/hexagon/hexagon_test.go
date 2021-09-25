package hexagon

import (
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
		want Vector2
	}{
		{"0", args{h: hexagon{
			column: 0,
			row:    0,
		}}, Vector2{0.8660254037844386, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetCenterCoordinates(tt.args.h); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCenterCoordinates() = %v, want %v", got, tt.want)
			}
		})
	}
}
