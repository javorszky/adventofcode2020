package day23

import (
	"container/ring"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_step(t *testing.T) {
	type args struct {
		inputs []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "correctly does my own step",
			args: args{
				inputs: []int{1, 3, 5, 4, 6, 8, 7, 2, 9},
			},
			want: []int{6, 8, 7, 2, 9, 3, 5, 4, 1},
		},
		{
			name: "correctly does step 1 from example",
			args: args{
				inputs: []int{3, 8, 9, 1, 2, 5, 4, 6, 7},
			},
			want: []int{2, 8, 9, 1, 5, 4, 6, 7, 3},
		},
		{
			name: "correctly does step 2 from example",
			args: args{
				inputs: []int{2, 8, 9, 1, 5, 4, 6, 7, 3},
			},
			want: []int{5, 4, 6, 7, 8, 9, 1, 3, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// create the list
			r := ring.New(len(tt.args.inputs))

			addresses := make(map[int]*ring.Ring, 0)
			for _, v := range tt.args.inputs {
				r.Value = cup{
					Value:     v,
					Addresses: &addresses,
				}
				addresses[v] = r
				r = r.Next()
			}

			nr := step(r)

			out := make([]int, 0)
			for i := 0; i < nr.Len(); i++ {
				out = append(out, nr.Value.(cup).Value)
				nr = nr.Next()
			}

			assert.Equal(t, tt.want, out)
		})
	}
}

func Test_task1GetLabels(t *testing.T) {
	type args struct {
		inputs []int
		steps  int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "parses example correctly",
			args: args{
				inputs: []int{3, 8, 9, 1, 2, 5, 4, 6, 7},
				steps:  10,
			},
			want: "92658374",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, task1GetLabel(tt.args.inputs, tt.args.steps))
		})
	}
}
