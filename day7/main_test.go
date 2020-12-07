package day7

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_t2accumulator(t *testing.T) {
	var m = map[string]map[string]int{
		"bright beige": {
			"drab red":      4,
			"dull silver":   4,
			"pale violet":   2,
			"vibrant brown": 5,
		},
		"drab red": {
			"vibrant brown": 2,
		},
		"vibrant brown": {
			"bright aqua": 1,
		},
		"pale violet": {
			"dull tan": 2,
		},
		"dull tan": {
			"faded cyan":   4,
			"light purple": 2,
			"shiny indigo": 4,
		},
	}

	type args struct {
		m map[string]map[string]int
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "successfully calculates contained bags for pale violet",
			args: args{
				m: m,
				s: "vibrant brown",
			},
			want: 1,
		},
		{
			name: "successfully calculates contained bags for vibrant brown",
			args: args{
				m: m,
				s: "pale violet",
			},
			want: 22,
		},
		{
			name: "successfully calculates contained bags for shiny indigo",
			args: args{
				m: m,
				s: "shiny indigo",
			},
			want: 0,
		},
		{
			name: "successfully calculates contained bags for drab red",
			args: args{
				m: m,
				s: "drab red",
			},
			want: 4,
		},
		{
			name: "successfully calculates contained bags for bright beige",
			args: args{
				m: m,
				s: "bright beige",
			},
			want: 80,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, t2accumulator(tt.args.m, tt.args.s))
		})
	}
}
