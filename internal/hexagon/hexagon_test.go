package hexagon

import (
	"github.com/CmdSoda/boardgamewars/internal/math"
	"github.com/CmdSoda/boardgamewars/internal/segment"
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

func TestGetSegments(t *testing.T) {
	type fields struct {
		column int
		row    int
	}
	tests := []struct {
		name   string
		fields fields
		want   [6]segment.Segment
	}{
		{
			name: "01",
			fields: fields{
				column: 0,
				row:    0,
			},
			want: [6]segment.Segment{{
				// NW
				Start: math.Vector2{
					X: 0,
					Y: 0.5,
				},
				End: math.Vector2{
					X: 0.8660254037844386,
					Y: 0,
				},
			}, {
				// NE
				Start: math.Vector2{
					X: 0.8660254037844386,
					Y: 0,
				},
				End: math.Vector2{
					X: 1.7320508075688772,
					Y: 0.5,
				},
			}, {
				// E
				Start: math.Vector2{
					X: 1.7320508075688772,
					Y: 0.5,
				},
				End: math.Vector2{
					X: 1.7320508075688772,
					Y: 1.5,
				},
			}, {
				// SE
				Start: math.Vector2{
					X: 1.7320508075688772,
					Y: 1.5,
				},
				End: math.Vector2{
					X: 0.8660254037844386,
					Y: 2,
				},
			}, {
				// SW
				Start: math.Vector2{
					X: 0.8660254037844386,
					Y: 2,
				},
				End: math.Vector2{
					X: 0,
					Y: 1.5,
				},
			}, {
				// W
				Start: math.Vector2{
					X: 0,
					Y: 1.5,
				},
				End: math.Vector2{
					X: 0,
					Y: 0.5,
				},
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := hexagon{
				column: tt.fields.column,
				row:    tt.fields.row,
			}
			if got := h.GetSegments(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSegments() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetAdjacent(t *testing.T) {
	type fields struct {
		column int
		row    int
	}
	type args struct {
		direction segment.Direction
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   hexagon
	}{
		{
			name: "NE",
			fields: fields{
				column: 0,
				row:    1,
			},
			args: args{direction: segment.NE},
			want: hexagon{
				column: 1,
				row:    0,
			},
		},
		{
			name: "E",
			fields: fields{
				column: 0,
				row:    0,
			},
			args: args{direction: segment.E},
			want: hexagon{
				column: 1,
				row:    0,
			},
		},
		{
			name: "SE",
			fields: fields{
				column: 0,
				row:    0,
			},
			args: args{direction: segment.SE},
			want: hexagon{
				column: 0,
				row:    1,
			},
		},
		{
			name: "SW",
			fields: fields{
				column: 1,
				row:    0,
			},
			args: args{direction: segment.SW},
			want: hexagon{
				column: 0,
				row:    1,
			},
		},
		{
			name: "W",
			fields: fields{
				column: 1,
				row:    0,
			},
			args: args{direction: segment.W},
			want: hexagon{
				column: 0,
				row:    0,
			},
		},
		{
			name: "NW",
			fields: fields{
				column: 0,
				row:    1,
			},
			args: args{direction: segment.NW},
			want: hexagon{
				column: 0,
				row:    0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := hexagon{
				column: tt.fields.column,
				row:    tt.fields.row,
			}
			if got := h.GetAdjacent(tt.args.direction); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAdjacent() = %v, want %v", got, tt.want)
			}
		})
	}
}
