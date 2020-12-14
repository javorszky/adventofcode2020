package day14

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_applyMaskT2(t *testing.T) {
	type args struct {
		mask   string
		binary string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "applies mask correctly",
			args: args{
				mask:   "10X10X",
				binary: "111000",
			},
			want: "11X10X",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, applyMaskT2(tt.args.mask, tt.args.binary))
		})
	}
}

func Test_extrapolateAddresses(t *testing.T) {
	type args struct {
		prefix string
		s      string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "extrapolates from string where X is at the beginning",
			args: args{
				prefix: "",
				s:      "X11010101",
			},
			want: []string{
				"011010101",
				"111010101",
			},
		},
		{
			name: "extrapolates from string where X is at the end",
			args: args{
				prefix: "",
				s:      "11010101X",
			},
			want: []string{
				"110101010",
				"110101011",
			},
		},
		{
			name: "extrapolates from string where there are 3X: beginning, mid, end",
			args: args{
				prefix: "",
				s:      "X101X101X",
			},
			want: []string{
				"010101010",
				"010101011",
				"010111010",
				"010111011",
				"110101010",
				"110101011",
				"110111010",
				"110111011",
			},
		},
		{
			name: "returns the string as is if there are no Xes in it",
			args: args{
				prefix: "",
				s:      "011010101",
			},
			want: []string{
				"011010101",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.ElementsMatch(t, tt.want, extrapolateAddresses(tt.args.prefix, tt.args.s))
		})
	}
}
