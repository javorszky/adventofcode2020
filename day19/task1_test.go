package day19

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseRules(t *testing.T) {
	type args struct {
		ruleStrings []string
	}
	tests := []struct {
		name string
		args args
		want map[string]string
	}{
		{
			name: "parses rule string slice correctly",
			args: args{
				ruleStrings: []string{
					"1: 1 2 | 3 4",
					"2: \"a\"",
				},
			},
			want: map[string]string{
				"1": "1 2 | 3 4",
				"2": "\"a\"",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, parseRules(tt.args.ruleStrings))
		})
	}
}

func Test_findRule(t *testing.T) {
	type args struct {
		s     string
		rules map[string]string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "finds rule where rule is 'a'",
			args: args{
				s: "0",
				rules: map[string]string{
					"0": `"a"`,
				},
			},
			want: `a`,
		},
		{
			name: "finds rule where rule is 'b'",
			args: args{
				s: "0",
				rules: map[string]string{
					"0": `"b"`,
				},
			},
			want: `b`,
		},
		{
			name: "finds rule where rule is 'a' through substitution",
			args: args{
				s: "0",
				rules: map[string]string{
					"0": "1",
					"1": `"a"`,
				},
			},
			want: `( a )`,
		},
		{
			name: "finds rule where rule is 'a' and 'b' through substitutions with an or clause ",
			args: args{
				s: "0",
				rules: map[string]string{
					"0": "1 | 2",
					"1": `"a"`,
					"2": `"b"`,
				},
			},
			want: `( a | b )`,
		},
		{
			name: "finds rule where rule is 'a' and 'b' through substitutions with an or clause and repetitions",
			args: args{
				s: "0",
				rules: map[string]string{
					"0": "1 1 | 2 2",
					"1": `"a"`,
					"2": `"b"`,
				},
			},
			want: `( a a | b b )`,
		},
		{
			name: "finds rule where rule is 'a' and 'b' through substitutions with an or clause and repetitions in multiple levels",
			args: args{
				s: "0",
				rules: map[string]string{
					"0": "1 1 | 2 2",
					"1": "3 4 | 4 3",
					"2": `"b"`,
					"3": "2 5",
					"4": "3 2 | 2 3",
					"5": `"a"`,
				},
			},
			want: `( ( ( b a ) ( ( b a ) b | b ( b a ) ) | ( ( b a ) b | b ( b a ) ) ( b a ) ) ( ( b a ) ( ( b a ) b | b ( b a ) ) | ( ( b a ) b | b ( b a ) ) ( b a ) ) | b b )`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, findRule(tt.args.s, tt.args.rules))
		})
	}
}
