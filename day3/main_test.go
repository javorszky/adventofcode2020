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
			want: true,
		},
		{
			name: "correctly returns tree in row 1 not overflown",
			args: args{
				row:    1,
				column: 4,
				slope:  []string{"#......##.", "....#..##."},
			},
			want: true,
		},
		{
			name: "correctly returns not-tree in row 1 not overflown",
			args: args{
				row:    1,
				column: 3,
				slope:  []string{"#......##.", "....#..##."},
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, isTree(tt.args.row, tt.args.column, tt.args.slope))
		})
	}
}
