package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isTree(t *testing.T) {
	type args struct {
		row    int
		column int
		slope  []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "correctly returns tree in row 0 not overflown",
			args: args{
				row:    0,
				column: 0,
				slope:  []string{"#......##."},
			},
			want: true,
		},
		{
			name: "correctly returns not-tree in row 0 not overflown",
			args: args{
				row:    0,
				column: 1,
				slope:  []string{"#......##."},
			},
			want: false,
		},
		{
			name: "correctly returns tree in row 0 overflown",
			args: args{
				row:    0,
				column: 10,
				slope:  []string{"#......##."},
			},
			want: true,
		},
		{
			name: "correctly returns not-tree in row 0 overflown",
			args: args{
				row:    0,
				column: 11,
				slope:  []string{"#......##."},
			},
			want: false,
		},
		{
			name: "correctly returns tree in row 1 not overflown",
			args: args{
				row:    1,
				column: 4,
				slope: []string{
					"#......##.",
					"....#..##.",
				},
			},
			want: true,
		},
		{
			name: "correctly returns not-tree in row 1 not overflown",
			args: args{
				row:    1,
				column: 3,
				slope: []string{
					"#......##.",
					"....#..##.",
				},
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, isTree(tt.args.row, tt.args.column, tt.args.slope))
		})
	}
}

func Test_step(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{
			name: "correctly gets result for step 1",
			args: args{
				n: 1,
			},
			want:  1,
			want1: 3,
		},
		{
			name: "correctly gets result for step 6",
			args: args{
				n: 6,
			},
			want:  6,
			want1: 18,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := step(tt.args.n)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want1, got1)
		})
	}
}
