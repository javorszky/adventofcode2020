package day23

import (
	"container/ring"
	"testing"
	"time"

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

func Test_generateRing(t *testing.T) {
	type args struct {
		start []int
		total int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "generates tail padded ring",
			args: args{
				start: []int{9, 3, 6, 1, 2, 4, 7, 8, 5},
				total: 15,
			},
			want: []int{9, 3, 6, 1, 2, 4, 7, 8, 5, 10, 11, 12, 13, 14, 15},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := generateRing(tt.args.start, tt.args.total)
			helper := make([]int, 0)
			for i := 0; i < r.Len(); i++ {
				helper = append(helper, r.Value.(cup).Value)
				r = r.Next()
			}
			assert.Equal(t, tt.want, helper)
		})
	}
}

func Test_task2GetProducts(t *testing.T) {
	type args struct {
		inputs    []int
		totalCups int
		steps     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "do the example task 2",
			args: args{
				inputs:    []int{3, 8, 9, 1, 2, 5, 4, 6, 7},
				totalCups: cupsInCircle,
				steps:     moveCupsThisManyTimes,
			},
			want: 149245887792,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// timeout implementation from https://erikwinter.nl/notes/2020/setting-a-timeout-on-golang-unit-tests/.
			timeout := time.After(4 * time.Second)
			done := make(chan bool)

			go func() {
				assert.Equal(t, tt.want, task2GetProducts(tt.args.inputs, tt.args.totalCups, tt.args.steps))

				done <- true
			}()

			select {
			case <-timeout:
				t.Fatal("test didn't finish in 4 seconds")
			case <-done:
			}
		})
	}
}
