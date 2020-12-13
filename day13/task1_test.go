package day13

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_t1formatInput(t *testing.T) {
	type args struct {
		in []string
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 []int
	}{
		{
			name: "turns input string into timecode and bus slices",
			args: args{
				in: []string{
					"18237",
					"1,2,3,x,x,x,4,x,645,x,x,x,4,x,x,x,x,9",
				},
			},
			want:  18237,
			want1: []int{1, 2, 3, 4, 645, 4, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := t1formatInput(tt.args.in)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want1, got1)
		})
	}
}

func Test_t1next(t *testing.T) {
	type args struct {
		t        int
		interval int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "returns 0 if t is divisible by interval",
			args: args{
				t:        95,
				interval: 5,
			},
			want: 0,
		},
		{
			name: "returns 4 if the next multiple of interval is 4 from t",
			args: args{
				t:        96,
				interval: 5,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, t1next(tt.args.t, tt.args.interval))
		})
	}
}
