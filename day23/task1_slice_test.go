package day23

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_createCircleSlice(t *testing.T) {
	type args struct {
		in []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "creates a linked slice",
			args: args{
				in: []int{3, 8, 9, 1, 2, 5, 4, 6, 7},
			}, //       0, 1, 2, 3, 4, 5, 6, 7, 8, 9
			want: []int{0, 2, 5, 8, 6, 4, 7, 3, 9, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, createCircleSlice(tt.args.in))
		})
	}
}

func Test_cupGame_play(t *testing.T) {
	type fields struct {
		Elements []int
		current  int
		max      int
	}
	tests := []struct {
		name        string
		fields      fields
		want        []int
		wantCurrent int
	}{
		{
			name: "plays a round correctly, aoc example, round 1",
			fields: fields{
				Elements: []int{0, 2, 5, 8, 6, 4, 7, 3, 9, 1},
				current:  3,
				max:      9,
			}, //              0  1  2  3  4  5  6  7  8  9
			want:        []int{0, 5, 8, 2, 6, 4, 7, 3, 9, 1},
			wantCurrent: 2,
		},
		{
			name: "plays a round correctly, aoc example, round 2",
			fields: fields{
				Elements: []int{0, 5, 8, 2, 6, 4, 7, 3, 9, 1},
				current:  2,
				max:      9,
			}, //              0  1  2  3  4  5  6  7  8  9
			want:        []int{0, 3, 5, 2, 6, 4, 7, 8, 9, 1},
			wantCurrent: 5,
		},
		{
			name: "plays a round correctly, aoc example, round 3",
			fields: fields{
				Elements: []int{0, 3, 5, 2, 6, 4, 7, 8, 9, 1},
				current:  5,
				max:      9,
			}, //              0  1  2  3  4  5  6  7  8  9
			want:        []int{0, 3, 5, 4, 6, 8, 7, 2, 9, 1},
			wantCurrent: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &cupGame{
				Elements: tt.fields.Elements,
				Current:  tt.fields.current,
				max:      tt.fields.max,
			}

			c.play()

			assert.Equal(t, tt.want, c.Elements)
			assert.Equal(t, tt.wantCurrent, c.Current)
		})
	}
}

func Test_getLabel(t *testing.T) {
	type args struct {
		in   []int
		step int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "gets label correctly for input and number of games",
			args: args{
				in:   []int{3, 8, 9, 1, 2, 5, 4, 6, 7},
				step: 100,
			},
			want: "67384529",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, getLabel(tt.args.in, tt.args.step))
		})
	}
}
