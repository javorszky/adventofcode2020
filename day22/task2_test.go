package day22

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_playString(t *testing.T) {
	type args struct {
		p1 []int
		p2 []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "converts two decks into a string",
			args: args{
				p1: []int{1, 2, 3, 4, 5},
				p2: []int{6, 7, 8, 9, 10},
			},
			want: "1,2,3,4,5|6,7,8,9,10",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, playString(tt.args.p1, tt.args.p2))
		})
	}
}

func Test_game(t *testing.T) {
	type args struct {
		p1Deck []int
		p2Deck []int
		level  int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 []int
		want2 []int
	}{
		{
			name: "plays recursive combat per aoc d22 t2 example",
			args: args{
				p1Deck: []int{9, 2, 6, 3, 1},
				p2Deck: []int{5, 8, 4, 7, 10},
				level:  1,
			},
			want:  2,
			want1: []int{},
			want2: []int{7, 5, 6, 2, 4, 1, 10, 8, 9, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := game(tt.args.p1Deck, tt.args.p2Deck, tt.args.level)
			if got != tt.want {
				t.Errorf("game() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("game() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("game() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}
