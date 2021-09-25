package path

import (
	"github.com/CmdSoda/boardgamewars/internal/hexagon"
	"reflect"
	"testing"
)

func TestCalculatePath(t *testing.T) {
	type args struct {
		startHex hexagon.Hexagon
		endHex   hexagon.Hexagon
	}
	tests := []struct {
		name string
		args args
		want []hexagon.Hexagon
	}{
		{
			name: "01",
			args: args{
				startHex: hexagon.NewHexagon(1, 1),
				endHex:   hexagon.NewHexagon(1, 3),
			},
			want: []hexagon.Hexagon{
				{Column: 1, Row: 1},
				{Column: 1, Row: 2},
				{Column: 1, Row: 3}},
		},
		{
			name: "02",
			args: args{
				startHex: hexagon.NewHexagon(1, 2),
				endHex:   hexagon.NewHexagon(7, 3),
			},
			want: []hexagon.Hexagon{
				{Column: 1, Row: 2},
				{Column: 2, Row: 2},
				{Column: 3, Row: 2},
				{Column: 4, Row: 2},
				{Column: 5, Row: 3},
				{Column: 6, Row: 3},
				{Column: 7, Row: 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculatePath(tt.args.startHex, tt.args.endHex); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CalculatePath() = %v, want %v", got, tt.want)
			}
		})
	}
}
