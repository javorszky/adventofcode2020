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
