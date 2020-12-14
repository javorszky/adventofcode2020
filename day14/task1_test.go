package day14

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_decimalToBinary(t *testing.T) {
	type args struct {
		i int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "converts decimal to binary string where value is all 1s",
			args: args{
				// 2^37-1, should overflow
				i: 137438953472 - 1,
			},
			want: "111111111111111111111111111111111111",
		},
		{
			name: "converts decimal to binary string where decimal overflows",
			args: args{
				// 2^41, should overflow
				i: 2199023255552,
			},
			want: "000000000000000000000000000000000000",
		},
		{
			name: "converts decimal to binary string",
			args: args{
				i: 1,
			},
			want: "000000000000000000000000000000000001",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, decimalToBinary(tt.args.i))
		})
	}
}

func Test_binaryToDecimal(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "converts all zero binary string to decimal",
			args: args{
				s: "000000000000000000000000000000000000",
			},
			want: 0,
		},
		{
			name: "converts all 1 binary string to decimal",
			args: args{
				s: "111111111111111111111111111111111111",
			},
			want: 68719476735,
		},
		{
			name: "converts too long all 1 binary string to decimal",
			args: args{
				s: "1111111111111111111111111111111111111111111111111",
			},
			want: 68719476735,
		},
		{
			name: "converts too long all 1 binary string to decimal cutting off correct end",
			args: args{
				s: "1100000000000000000000000000000000000000000000011",
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, binaryToDecimal(tt.args.s))
		})
	}
}

func Test_applyMask(t *testing.T) {
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
			want: "101100",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, applyMask(tt.args.mask, tt.args.binary))
		})
	}
}
