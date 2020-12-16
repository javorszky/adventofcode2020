package day16

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_filterInvalidTickets(t *testing.T) {
	type args struct {
		values        map[int]struct{}
		nearbyTickets []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "filters out invalid tickets",
			args: args{
				values: map[int]struct{}{
					23: {},
				},
				nearbyTickets: []string{
					"12,13,14,15",
					"16,17,18,19",
					"20,21,22,23",
					"24,25,26,27",
				},
			},
			want: []string{
				"20,21,22,23",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, filterInvalidTickets(tt.args.values, tt.args.nearbyTickets))
		})
	}
}

func Test_newRule(t *testing.T) {
	type args struct {
		name string
		n1   string
		n2   string
	}
	tests := []struct {
		name string
		args args
		want rule
	}{
		{
			name: "creates a rule based on data in correct order",
			args: args{
				name: "some rule",
				n1:   "1",
				n2:   "2",
			},
			want: rule{
				name: "some rule",
				min:  1,
				max:  2,
			},
		},
		{
			name: "creates a rule based on data in reverse order",
			args: args{
				name: "some rule",
				n1:   "2",
				n2:   "1",
			},
			want: rule{
				name: "some rule",
				min:  1,
				max:  2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, newRule(tt.args.name, tt.args.n1, tt.args.n2))
		})
	}
}

func Test_createRules(t *testing.T) {
	type args struct {
		rules []string
	}
	tests := []struct {
		name string
		args args
		want []rule
	}{
		{
			name: "creates rules for slice of strings",
			args: args{
				rules: []string{
					"departure location: 35-796 or 811-953",
					"departure station: 25-224 or 248-952",
				},
			},
			want: []rule{
				{
					name: "departure location",
					min:  35,
					max:  796,
				},
				{
					name: "departure location",
					min:  811,
					max:  953,
				},
				{
					name: "departure station",
					min:  25,
					max:  224,
				},
				{
					name: "departure station",
					min:  248,
					max:  952,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, createRules(tt.args.rules))
		})
	}
}
