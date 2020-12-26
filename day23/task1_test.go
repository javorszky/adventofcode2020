package day23

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_rotateCircleBy(t *testing.T) {
	type args struct {
		in []int
		n  int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "rotates an array by 1",
			args: args{
				in: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0},
				n:  1,
			},
			want: []int{2, 3, 4, 5, 6, 7, 8, 9, 0, 1},
		},
		{
			name: "rotates an array by 0",
			args: args{
				in: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0},
				n:  0,
			},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0},
		},
		{
			name: "rotates an array by 4",
			args: args{
				in: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0},
				n:  4,
			},
			want: []int{5, 6, 7, 8, 9, 0, 1, 2, 3, 4},
		},
		{
			name: "rotates an array by 12",
			args: args{
				in: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0},
				n:  12,
			},
			want: []int{3, 4, 5, 6, 7, 8, 9, 0, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, rotateCircleBy(tt.args.in, tt.args.n))
		})
	}
}

func Test_removeThreeCups(t *testing.T) {
	type args struct {
		in []int
	}
	tests := []struct {
		name  string
		args  args
		want  []int
		want1 []int
	}{
		{
			name: "removes three cups from the circle where the idx 0 is the Current one",
			args: args{
				in: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0},
			},
			want:  []int{1, 5, 6, 7, 8, 9, 0},
			want1: []int{2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := removeThreeCups(tt.args.in)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want1, got1)
		})
	}
}

func Test_findIndexOfNextCurrentCup(t *testing.T) {
	type args struct {
		in []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "finds the next Current cup when all elements are present",
			args: args{
				in: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0},
			},
			want: 9,
		},
		{
			name: "finds the next Current cup when all elements are present, order is random",
			args: args{
				in: []int{1, 9, 5, 7, 3, 0, 2, 4, 8, 6},
			},
			want: 5,
		},
		{
			name: "finds the next Current cup when some elements are missing, order is random",
			args: args{
				in: []int{1, 5, 7, 3, 2, 4, 6},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, findIndexOfNextCurrentCup(tt.args.in))
		})
	}
}

func Test_insertSnippetIntoAt(t *testing.T) {
	type args struct {
		short   []int
		snippet []int
		idx     int
	}
	tests := []struct {
		name       string
		args       args
		want       []int
		wantsPanic bool
	}{
		{
			name: "inserts slice into other slice at index 0",
			args: args{
				short:   []int{1, 2, 3, 4, 5, 6, 7},
				snippet: []int{8, 9, 0},
				idx:     0,
			},
			want:       []int{8, 9, 0, 1, 2, 3, 4, 5, 6, 7},
			wantsPanic: false,
		},
		{
			name: "inserts slice into other slice at index 4",
			args: args{
				short:   []int{1, 2, 3, 4, 5, 6, 7},
				snippet: []int{8, 9, 0},
				idx:     4,
			},
			want:       []int{1, 2, 3, 4, 8, 9, 0, 5, 6, 7},
			wantsPanic: false,
		},
		{
			name: "inserts slice into other slice at index 6",
			args: args{
				short:   []int{1, 2, 3, 4, 5, 6, 7},
				snippet: []int{8, 9, 0},
				idx:     6,
			},
			want:       []int{1, 2, 3, 4, 5, 6, 8, 9, 0, 7},
			wantsPanic: false,
		},
		{
			name: "inserts slice into other slice at index 7",
			args: args{
				short:   []int{1, 2, 3, 4, 5, 6, 7},
				snippet: []int{8, 9, 0},
				idx:     7,
			},
			want:       []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0},
			wantsPanic: false,
		},
		{
			name: "panics when trying to insert at index higher than the length of the original slice",
			args: args{
				short:   []int{1, 2, 3, 4, 5, 6, 7},
				snippet: []int{8, 9, 0},
				idx:     19,
			},
			want:       nil,
			wantsPanic: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantsPanic {
				assert.Panics(t, func() {
					insertSnippetIntoAt(tt.args.short, tt.args.snippet, tt.args.idx)
				})
				return
			}
			assert.Equal(t, tt.want, insertSnippetIntoAt(tt.args.short, tt.args.snippet, tt.args.idx))
		})
	}
}

func Test_round(t *testing.T) {
	type args struct {
		in []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "plays a round (round 1 example) correctly",
			args: args{
				in: []int{3, 8, 9, 1, 2, 5, 4, 6, 7},
			},
			want: []int{2, 8, 9, 1, 5, 4, 6, 7, 3},
		},
		{
			name: "plays a round (round 2 example) correctly",
			args: args{
				in: []int{2, 8, 9, 1, 5, 4, 6, 7, 3},
			},
			want: []int{5, 4, 6, 7, 8, 9, 1, 3, 2},
		},
		{
			name: "plays a round correctly",
			args: args{
				in: []int{5, 4, 6, 7, 8, 9, 1, 3, 2},
			},
			want: []int{8, 9, 1, 3, 4, 6, 7, 2, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, round(tt.args.in))
		})
	}
}

func Test_orderOfCups(t *testing.T) {
	type args struct {
		in []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "gets correct order of cups from slice",
			args: args{
				in: []int{5, 8, 3, 7, 4, 1, 9, 2, 6},
			},
			want: "92658374",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, orderOfCups(tt.args.in))
		})
	}
}
