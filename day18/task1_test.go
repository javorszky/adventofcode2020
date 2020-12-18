package day18

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_evaluateExpression(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			// pass
			name: "evaluates '9+4' correctly",
			args: args{
				s: "9+4",
			},
			want: 13,
		},
		{
			// pass
			name: "evaluates '9*4' correctly",
			args: args{
				s: "9*4",
			},
			want: 36,
		},
		{
			// fail
			name: "evaluates '9+4*5' correctly",
			args: args{
				s: "9+4*5",
			},
			want: 65,
		},
		{
			name: "evaluates '9*4+5' correctly",
			args: args{
				s: "9*4+5",
			},
			want: 41,
		},
		{
			name: "evaluates '9+4+5' correctly",
			args: args{
				s: "9+4+5",
			},
			want: 18,
		},
		{
			name: "evaluates '9+4+5+2' correctly",
			args: args{
				s: "9+4+5+2",
			},
			want: 20,
		},
		{
			name: "evaluates '9+4*5+2' correctly",
			args: args{
				s: "9+4*5+2",
			},
			want: 67,
		},
		{
			// fail, get 29, should be 31
			name: "evaluates '9+(4*5)+2' correctly",
			args: args{
				s: "9+(4*5)+2",
			},
			want: 31,
		},
		{
			name: "evaluates '9+4*(5+2)' correctly",
			args: args{
				s: "9+4*(5+2)",
			},
			want: 91,
		},
		{
			// fail, get 13, want 91 (13)*(7)
			name: "evaluates '(9+4)*(5+2)' correctly",
			args: args{
				s: "(9+4)*(5+2)",
			},
			want: 91,
		},
		{
			// pass though
			name: "evaluates '9+(4*(5+2))' correctly",
			args: args{
				s: "9+(4*(5+2))",
			},
			want: 37,
		},
		{
			// pass though
			name: "aoc example 1: '2*3+(4*5)'",
			args: args{
				s: "2*3+(4*5)",
			},
			want: 26,
		},
		{
			// pass though
			name: "aoc example 2: '5+(8*3+9+3*4*3)'",
			args: args{
				s: "5+(8*3+9+3*4*3)",
			},
			want: 437,
		},
		{
			// pass though
			name: "aoc example 3: '5*9*(7*3*3+9*3+(8+6*4))'",
			args: args{
				s: "5*9*(7*3*3+9*3+(8+6*4))",
			},
			want: 12240,
		},
		{
			// pass though
			name: "aoc example 4: '((2+4*9)*(6+9*8+6)+6)+2+4*2'",
			args: args{
				s: "((2+4*9)*(6+9*8+6)+6)+2+4*2",
			},
			want: 13632,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, evaluateExpression(tt.args.s))
		})
	}
}
