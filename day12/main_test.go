package day12

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_newDirection(t *testing.T) {
	type args struct {
		base string
		amnt int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "east plus 90 produces south",
			args: args{
				base: east,
				amnt: 90,
			},
			want: south,
		},
		{
			name: "east plus 180 produces west",
			args: args{
				base: east,
				amnt: 180,
			},
			want: west,
		},
		{
			name: "east plus 270 produces north",
			args: args{
				base: east,
				amnt: 270,
			},
			want: north,
		},
		{
			name: "south plus 90 produces west",
			args: args{
				base: south,
				amnt: 90,
			},
			want: west,
		},
		{
			name: "south plus 180 produces north",
			args: args{
				base: south,
				amnt: 180,
			},
			want: north,
		},
		{
			name: "south plus 270 produces east",
			args: args{
				base: south,
				amnt: 270,
			},
			want: east,
		},
		{
			name: "west plus 90 produces north",
			args: args{
				base: west,
				amnt: 90,
			},
			want: north,
		},
		{
			name: "west plus 180 produces east",
			args: args{
				base: west,
				amnt: 180,
			},
			want: east,
		},
		{
			name: "west plus 270 produces south",
			args: args{
				base: west,
				amnt: 270,
			},
			want: south,
		},
		{
			name: "north plus 90 produces east",
			args: args{
				base: north,
				amnt: 90,
			},
			want: east,
		},
		{
			name: "north plus 180 produces south",
			args: args{
				base: north,
				amnt: 180,
			},
			want: south,
		},
		{
			name: "north plus 270 produces west",
			args: args{
				base: north,
				amnt: 270,
			},
			want: west,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, newDirection(tt.args.base, tt.args.amnt))
		})
	}
}
