package day22

import (
	"reflect"
	"testing"
)

func Test_parseDeck(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "parses deck string into int",
			args: args{
				s: `Player 1:
9
2
6
3
1`,
			},
			want: []int{9, 2, 6, 3, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseDeck(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseDeck() = %v, want %v", got, tt.want)
			}
		})
	}
}
