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
			name: "converts decimal to binary string",
			args: args{
				// 2^37-1, should overflow
				i: 137438953472 - 1,
			},
			want: "111111111111111111111111111111111111",
		},
		{
			name: "converts decimal to binary string",
			args: args{
				// 2^41, should overflow
				i: 2199023255552,
			},
			want: "000000000000000000000000000000000000",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, decimalToBinary(tt.args.i))
		})
	}
}
