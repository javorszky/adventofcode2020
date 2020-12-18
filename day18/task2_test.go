package day18

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_evaluateExpression2(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "evaluates '9+4",
			args: args{
				s: "9+4",
			},
			want: 13,
		},
		{
			name: "evaluates '5*9+4",
			args: args{
				s: "5*9+4",
			},
			want: 65,
		},
		{
			name: "evaluates '5+9*4",
			args: args{
				s: "5+9*4",
			},
			want: 56,
		},
		{
			name: "aoc t2 example 1: '1+(2*3)+(4*(5+6))'",
			args: args{
				s: "1+(2*3)+(4*(5+6))",
			},
			want: 51,
		},
		{
			name: "aoc t2 example 2: '2*3+(4*5)'",
			args: args{
				s: "2*3+(4*5)",
			},
			want: 46,
		},
		{
			name: "aoc t2 example 3: '5+(8*3+9+3*4*3)'",
			args: args{
				s: "5+(8*3+9+3*4*3)",
			},
			want: 1445,
		},
		{
			name: "aoc t2 example 4: '5*9*(7*3*3+9*3+(8+6*4))'",
			args: args{
				s: "5*9*(7*3*3+9*3+(8+6*4))",
			},
			want: 669060,
		},
		{
			name: "aoc t2 example 5: '((2+4*9)*(6+9*8+6)+6)+2+4*2'",
			args: args{
				s: "((2+4*9)*(6+9*8+6)+6)+2+4*2",
			},
			want: 23340,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, evaluateExpression2(tt.args.s))
		})
	}
}
