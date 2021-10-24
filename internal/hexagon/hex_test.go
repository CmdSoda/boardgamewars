package hexagon

import (
	"reflect"
	"testing"
)

func TestGetAdjacent(t *testing.T) {
	type fields struct {
		Column int
		Row    int
	}
	type args struct {
		direction Direction
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Position
	}{
		{
			name: "NE",
			fields: fields{
				Column: 2,
				Row:    2,
			},
			args: args{direction: NE},
			want: &Position{
				Column: 3,
				Row:    3,
			},
		},
		{
			name: "E",
			fields: fields{
				Column: 2,
				Row:    2,
			},
			args: args{direction: E},
			want: &Position{
				Column: 3,
				Row:    2,
			},
		},
		{
			name: "SE",
			fields: fields{
				Column: 2,
				Row:    2,
			},
			args: args{direction: SE},
			want: &Position{
				Column: 3,
				Row:    1,
			},
		},
		{
			name: "SW",
			fields: fields{
				Column: 2,
				Row:    2,
			},
			args: args{direction: SW},
			want: &Position{
				Column: 2,
				Row:    1,
			},
		},
		{
			name: "W",
			fields: fields{
				Column: 2,
				Row:    2,
			},
			args: args{direction: W},
			want: &Position{
				Column: 1,
				Row:    2,
			},
		},
		{
			name: "NW",
			fields: fields{
				Column: 2,
				Row:    2,
			},
			args: args{direction: NW},
			want: &Position{
				Column: 2,
				Row:    3,
			},
		},
		{
			name: "nil-1",
			fields: fields{
				Column: 1,
				Row:    1,
			},
			args: args{direction: NW},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := Position{
				Column: tt.fields.Column,
				Row:    tt.fields.Row,
			}
			if got := h.GetAdjacent(tt.args.direction); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAdjacent() = %v, want %v", got, tt.want)
			}
		})
	}
}
