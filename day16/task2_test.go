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

func Test_parseTicketPossibilities(t *testing.T) {
	type args struct {
		ticket string
		rules  []rule
	}
	tests := []struct {
		name string
		args args
		want map[int][]string
	}{
		{
			name: "parses possibilities correctly for ticket based on rules",
			args: args{
				ticket: "12,15,78,99",
				rules: []rule{
					{
						name: "rule1",
						min:  1,
						max:  78,
					},
					{
						name: "rule1",
						min:  80,
						max:  110,
					},
					{
						name: "rule2",
						min:  16,
						max:  78,
					},
					{
						name: "rule3",
						min:  79,
						max:  108,
					},
					{
						name: "rule4",
						min:  15,
						max:  61,
					},
				},
			},
			want: map[int][]string{
				1: {"rule1"},
				2: {"rule1", "rule4"},
				3: {"rule1", "rule2"},
				4: {"rule1", "rule3"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, parseTicketPossibilities(tt.args.ticket, tt.args.rules))
		})
	}
}
