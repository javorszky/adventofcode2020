package day16

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_mergeRanges(t *testing.T) {
	type args struct {
		values map[int]struct{}
		min    string
		max    string
	}
	tests := []struct {
		name string
		args args
		want map[int]struct{}
	}{
		{
			name: "adds entire range including min max to empty range",
			args: args{
				values: map[int]struct{}{},
				min:    "67",
				max:    "71",
			},
			want: map[int]struct{}{
				67: {},
				68: {},
				69: {},
				70: {},
				71: {},
			},
		},
		{
			name: "adds entire range including min max to empty range if min max are other way around",
			args: args{
				values: map[int]struct{}{},
				min:    "71",
				max:    "67",
			},
			want: map[int]struct{}{
				67: {},
				68: {},
				69: {},
				70: {},
				71: {},
			},
		},
		{
			name: "adds entire range including min max to non empty range where ranges overlap",
			args: args{
				values: map[int]struct{}{
					65: {},
					66: {},
					67: {},
					68: {},
					69: {},
				},
				min: "67",
				max: "71",
			},
			want: map[int]struct{}{
				65: {},
				66: {},
				67: {},
				68: {},
				69: {},
				70: {},
				71: {},
			},
		},
		{
			name: "adds entire range including min max to non empty range where ranges do not",
			args: args{
				values: map[int]struct{}{
					41: {},
					42: {},
					43: {},
				},
				min: "67",
				max: "71",
			},
			want: map[int]struct{}{
				41: {},
				42: {},
				43: {},
				67: {},
				68: {},
				69: {},
				70: {},
				71: {},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, mergeRanges(tt.args.values, tt.args.min, tt.args.max))
		})
	}
}
