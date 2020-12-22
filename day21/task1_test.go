package day21

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseInputs(t *testing.T) {
	type args struct {
		list []string
	}
	tests := []struct {
		name string
		args args
		want map[string][][]string
	}{
		{
			name: "parses input correctly",
			args: args{
				list: []string{
					"mxmxvkd kfcds sqjhc nhms (contains dairy, fish)",
					"trh fvjkl sbzzf mxmxvkd (contains dairy)",
					"sqjhc fvjkl (contains soy)",
					"sqjhc mxmxvkd sbzzf (contains fish)",
				},
			},
			want: map[string][][]string{
				"soy": {
					{
						"sqjhc",
						"fvjkl",
					},
				},
				"fish": {
					{
						"mxmxvkd",
						"kfcds",
						"sqjhc",
						"nhms",
					},
					{
						"sqjhc",
						"mxmxvkd",
						"sbzzf",
					},
				},
				"dairy": {
					{
						"mxmxvkd",
						"kfcds",
						"sqjhc",
						"nhms",
					},
					{
						"trh",
						"fvjkl",
						"sbzzf",
						"mxmxvkd",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, parseInputs(tt.args.list))
		})
	}
}

func Test_commonElements(t *testing.T) {
	type args struct {
		a []string
		b []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "returns common elements from both slices",
			args: args{
				a: []string{
					"a", "b", "c", "d",
				},
				b: []string{
					"a", "c", "e", "f", "q",
				},
			},
			want: []string{
				"a", "c",
			},
		},
		{
			name: "returns common elements from both slices if the slices are reversed",
			args: args{
				b: []string{
					"a", "b", "c", "d",
				},
				a: []string{
					"a", "c", "e", "f", "q",
				},
			},
			want: []string{
				"a", "c",
			},
		},
		{
			name: "returns common elements from both slices where there are no common elements",
			args: args{
				b: []string{
					"a", "b", "c", "d",
				},
				a: []string{
					"x", "y", "e", "f", "q",
				},
			},
			want: []string{},
		},
		{
			name: "returns common elements from both slices where one slice completely contains the other",
			args: args{
				b: []string{
					"a", "b", "c", "d",
				},
				a: []string{
					"a", "b", "c", "d", "x", "y", "e", "f", "q",
				},
			},
			want: []string{
				"a", "b", "c", "d",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, commonElements(tt.args.a, tt.args.b))
		})
	}
}

func Test_reduceElementsToCommon(t *testing.T) {
	type args struct {
		slices [][]string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "reduces slice of slices to their commone elements",
			args: args{
				slices: [][]string{
					{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"},
					{"a", "c", "e", "g", "i", "j"},
					{"n", "m", "l", "k", "j", "i"},
				},
			},
			want: []string{
				"j", "i",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, reduceElementsToCommon(tt.args.slices...))
		})
	}
}
