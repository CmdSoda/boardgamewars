package path

import (
	"github.com/CmdSoda/boardgamewars/internal/hexagon"
	"github.com/stretchr/testify/assert"
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
			name: "11->13",
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
			name: "12->73",
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
		{
			name: "11->73",
			args: args{
				startHex: hexagon.NewHexagon(1, 1),
				endHex:   hexagon.NewHexagon(7, 3),
			},
			want: []hexagon.Hexagon{
				{Column: 1, Row: 1},
				{Column: 2, Row: 1},
				{Column: 2, Row: 2},
				{Column: 3, Row: 2},
				{Column: 4, Row: 2},
				{Column: 5, Row: 2},
				{Column: 6, Row: 3},
				{Column: 7, Row: 3},
			},
		},
		{
			name: "23->21",
			args: args{
				startHex: hexagon.NewHexagon(2, 3),
				endHex:   hexagon.NewHexagon(2, 1),
			},
			want: []hexagon.Hexagon{
				{Column: 2, Row: 3},
				{Column: 2, Row: 2},
				{Column: 2, Row: 1},
			},
		},
		{
			name: "13->31",
			args: args{
				startHex: hexagon.NewHexagon(1, 3),
				endHex:   hexagon.NewHexagon(3, 1),
			},
			want: []hexagon.Hexagon{
				{Column: 1, Row: 3},
				{Column: 1, Row: 2},
				{Column: 2, Row: 2},
				{Column: 3, Row: 1},
			},
		},
		{
			name: "24->11",
			args: args{
				startHex: hexagon.NewHexagon(2, 4),
				endHex:   hexagon.NewHexagon(1, 1),
			},
			want: []hexagon.Hexagon{
				{Column: 2, Row: 4},
				{Column: 2, Row: 3},
				{Column: 1, Row: 2},
				{Column: 1, Row: 1},
			},
		},
		{
			name: "11->51",
			args: args{
				startHex: hexagon.NewHexagon(1, 1),
				endHex:   hexagon.NewHexagon(5, 1),
			},
			want: []hexagon.Hexagon{
				{Column: 1, Row: 1},
				{Column: 2, Row: 1},
				{Column: 3, Row: 1},
				{Column: 4, Row: 1},
				{Column: 5, Row: 1},
			},
		},
		{
			name: "12->52",
			args: args{
				startHex: hexagon.NewHexagon(1, 2),
				endHex:   hexagon.NewHexagon(5, 2),
			},
			want: []hexagon.Hexagon{
				{Column: 1, Row: 2},
				{Column: 2, Row: 2},
				{Column: 3, Row: 2},
				{Column: 4, Row: 2},
				{Column: 5, Row: 2},
			},
		},
		{
			name: "73->12",
			args: args{
				startHex: hexagon.NewHexagon(7, 3),
				endHex:   hexagon.NewHexagon(1, 2),
			},
			want: []hexagon.Hexagon{
				{Column: 7, Row: 3},
				{Column: 6, Row: 3},
				{Column: 5, Row: 3},
				{Column: 4, Row: 2},
				{Column: 3, Row: 2},
				{Column: 2, Row: 2},
				{Column: 1, Row: 2},
			},
		},
		{
			name: "41->13",
			args: args{
				startHex: hexagon.NewHexagon(4, 1),
				endHex:   hexagon.NewHexagon(1, 3),
			},
			want: []hexagon.Hexagon{
				{Column: 4, Row: 1},
				{Column: 3, Row: 2},
				{Column: 2, Row: 2},
				{Column: 2, Row: 3},
				{Column: 1, Row: 3},
			},
		},
		{
			name: "11->54",
			args: args{
				startHex: hexagon.NewHexagon(1, 1),
				endHex:   hexagon.NewHexagon(5, 4),
			},
			want: []hexagon.Hexagon{
				{Column: 1, Row: 1},
				{Column: 1, Row: 2},
				{Column: 2, Row: 2},
				{Column: 3, Row: 3},
				{Column: 4, Row: 3},
				{Column: 4, Row: 4},
				{Column: 5, Row: 4},
			},
		},
		{
			name: "82->84",
			args: args{
				startHex: hexagon.NewHexagon(8, 2),
				endHex:   hexagon.NewHexagon(8, 4),
			},
			want: []hexagon.Hexagon{
				{Column: 8, Row: 2},
				{Column: 9, Row: 3},
				{Column: 8, Row: 4},
			},
		},
		{
			name: "41->52",
			args: args{
				startHex: hexagon.NewHexagon(4, 1),
				endHex:   hexagon.NewHexagon(5, 2),
			},
			want: []hexagon.Hexagon{
				{Column: 4, Row: 1},
				{Column: 4, Row: 2},
				{Column: 5, Row: 2},
			},
		},
		{
			name: "41->73",
			args: args{
				startHex: hexagon.NewHexagon(4, 1),
				endHex:   hexagon.NewHexagon(7, 3),
			},
			want: []hexagon.Hexagon{
				{Column: 4, Row: 1},
				{Column: 4, Row: 2},
				{Column: 5, Row: 2},
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

const combinationMax int = 20

func TestCombination(t *testing.T) {
	for rowStart := 1; rowStart < combinationMax; rowStart++ {
		for columnStart := 1; columnStart < combinationMax; columnStart++ {
			for rowEnd := 1; rowEnd < combinationMax; rowEnd++ {
				for columnEnd := 1; columnEnd < combinationMax; columnEnd++ {
					path := CalculatePath(hexagon.NewHexagon(columnStart, rowStart), hexagon.NewHexagon(columnEnd, rowEnd))
					assert.Greater(t, len(path), 0)
				}
			}
		}
	}
}
