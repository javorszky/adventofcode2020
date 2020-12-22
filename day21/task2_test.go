package day21

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_removeElementFromSlice(t *testing.T) {
	type args struct {
		s     string
		slice []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "removing non existent element will return same slice",
			args: args{
				s:     "hello",
				slice: []string{"there", "general", "kenobi"},
			},
			want: []string{"there", "general", "kenobi"},
		},
		{
			name: "removes element from the beginning",
			args: args{
				s:     "hello",
				slice: []string{"hello", "there", "general", "kenobi"},
			},
			want: []string{"there", "general", "kenobi"},
		},
		{
			name: "removes element from the end",
			args: args{
				s:     "hello",
				slice: []string{"there", "general", "kenobi", "hello"},
			},
			want: []string{"there", "general", "kenobi"},
		},
		{
			name: "removes element from the middle",
			args: args{
				s:     "hello",
				slice: []string{"there", "hello", "general", "kenobi"},
			},
			want: []string{"there", "general", "kenobi"},
		},
		{
			name: "when multiple elements are present, only removes the first one",
			args: args{
				s:     "hello",
				slice: []string{"hello", "there", "hello", "general", "kenobi", "hello"},
			},
			want: []string{"there", "hello", "general", "kenobi", "hello"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.ElementsMatch(t, tt.want, removeElementFromSlice(tt.args.s, tt.args.slice))
		})
	}
}

func Test_sherlockFoodMap(t *testing.T) {
	type args struct {
		allergenFoodMap map[string][]string
	}
	tests := []struct {
		name string
		args args
		want map[string]string
	}{
		{
			name: "finds allergen food mapping",
			args: args{
				allergenFoodMap: map[string][]string{
					"dairy": {
						"mxmxvkd",
					},
					"fish": {
						"mxmxvkd",
						"sqjhc",
					},
					"soy": {
						"sqjhc",
						"fvjkl",
					},
				},
			},
			want: map[string]string{
				"dairy": "mxmxvkd",
				"fish":  "sqjhc",
				"soy":   "fvjkl",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, sherlockFoodMap(tt.args.allergenFoodMap))
		})
	}
}

func Test_foodMapping(t *testing.T) {
	type args struct {
		s []string
	}
	tests := []struct {
		name string
		args args
		want map[string]string
	}{
		{
			name: "generates allergen food map from input",
			args: args{
				s: []string{
					"mxmxvkd kfcds sqjhc nhms (contains dairy, fish)",
					"trh fvjkl sbzzf mxmxvkd (contains dairy)",
					"sqjhc fvjkl (contains soy)",
					"sqjhc mxmxvkd sbzzf (contains fish)",
				},
			},
			want: map[string]string{
				"dairy": "mxmxvkd",
				"fish":  "sqjhc",
				"soy":   "fvjkl",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, foodMapping(tt.args.s))
		})
	}
}

func Test_canonicalFoodList(t *testing.T) {
	type args struct {
		in map[string]string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "creates list of ingredients alphabetically sorted by the allergens",
			args: args{
				in: map[string]string{
					"dairy": "mxmxvkd",
					"fish":  "sqjhc",
					"soy":   "fvjkl",
				},
			},
			want: "mxmxvkd,sqjhc,fvjkl",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, canonicalFoodList(tt.args.in))
		})
	}
}
