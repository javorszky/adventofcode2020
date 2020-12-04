package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_identity_IsValid(t *testing.T) {
	type fields struct {
		ECL string
		PID string
		EYR string
		HCL string
		BYR string
		IYR string
		CID string
		HGT string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "valid struct with everything correct is valid",
			fields: fields{
				ECL: "amb",
				PID: "028372993",
				EYR: "2023",
				HCL: "#992233",
				BYR: "1984",
				IYR: "2013",
				CID: "gb",
				HGT: "182cm",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := identity{
				ECL: tt.fields.ECL,
				PID: tt.fields.PID,
				EYR: tt.fields.EYR,
				HCL: tt.fields.HCL,
				BYR: tt.fields.BYR,
				IYR: tt.fields.IYR,
				CID: tt.fields.CID,
				HGT: tt.fields.HGT,
			}

			assert.Equal(t, tt.want, i.IsValid())
		})
	}
}

func Test_identity_eclValid(t *testing.T) {
	type fields struct {
		ECL string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name:   "ecl is valid with amb",
			fields: fields{ECL: "amb"},
			want:   true,
		},
		{
			name:   "ecl is valid with blu",
			fields: fields{ECL: "blu"},
			want:   true,
		},
		{
			name:   "ecl is valid with brn",
			fields: fields{ECL: "brn"},
			want:   true,
		},
		{
			name:   "ecl is valid with gry",
			fields: fields{ECL: "gry"},
			want:   true,
		},
		{
			name:   "ecl is valid with grn",
			fields: fields{ECL: "grn"},
			want:   true,
		},
		{
			name:   "ecl is valid with hzl",
			fields: fields{ECL: "hzl"},
			want:   true,
		},
		{
			name:   "ecl is valid with oth",
			fields: fields{ECL: "oth"},
			want:   true,
		},
		{
			name:   "ecl is not valid with empty",
			fields: fields{ECL: ""},
			want:   false,
		},
		{
			name:   "ecl is not valid with random",
			fields: fields{ECL: "ylw"},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := identity{
				ECL: tt.fields.ECL,
			}
			assert.Equal(t, tt.want, i.eclValid())
		})
	}
}

func Test_identity_pidValid(t *testing.T) {

	tests := []struct {
		name string
		pid  string
		want bool
	}{
		{
			name: "pid is valid when it's 9 digits",
			pid:  "092837652",
			want: true,
		},
		{
			name: "pid is not valid when it's empty",
			pid:  "",
			want: false,
		},
		{
			name: "pid is not valid when it's 8 digits",
			pid:  "09283765",
			want: false,
		},
		{
			name: "pid is not valid when it's 10 digits",
			pid:  "0928376225",
			want: false,
		},
		{
			name: "pid is not valid when it's 9 non-digits",
			pid:  "abcdefghi",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := identity{
				PID: tt.pid,
			}

			assert.Equal(t, tt.want, i.pidValid())
		})
	}
}

func Test_identity_eyrValid(t *testing.T) {
	tests := []struct {
		name string
		eyr  string
		want bool
	}{
		{
			name: "lower bound valid eyr is valid",
			eyr:  "2020",
			want: true,
		},
		{
			name: "upper bound valid eyr is valid",
			eyr:  "2030",
			want: true,
		},
		{
			name: "midrange valid eyr is valid",
			eyr:  "2025",
			want: true,
		},
		{
			name: "under lower bound invalid eyr is invalid",
			eyr:  "2019",
			want: false,
		},
		{
			name: "above upper bound invalid eyr is invalid",
			eyr:  "2031",
			want: false,
		},
		{
			name: "empty string eyr is invalid",
			eyr:  "",
			want: false,
		},
		{
			name: "string eyr is invalid",
			eyr:  "boo",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := identity{
				EYR: tt.eyr,
			}

			assert.Equal(t, tt.want, i.eyrValid())
		})
	}
}

func Test_identity_hclValid(t *testing.T) {
	type fields struct {
		HCL string
	}
	tests := []struct {
		name string
		hcl  string
		want bool
	}{
		{
			name: "valid only digit hair color is valid",
			hcl:  "#123456",
			want: true,
		},
		{
			name: "valid 0-f hair color is valid",
			hcl:  "#07a7fb",
			want: true,
		},
		{
			name: "out of bounds hex color is out of bounds",
			hcl:  "#0123gh",
			want: false,
		},
		{
			name: "not starting with hash is invalid",
			hcl:  "123456",
			want: false,
		},
		{
			name: "hair colour with more than 6 characters is invalid",
			hcl:  "#123456abcdef",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := identity{
				HCL: tt.hcl,
			}
			assert.Equal(t, tt.want, i.hclValid())
		})
	}
}

func Test_identity_byrValid(t *testing.T) {
	tests := []struct {
		name string
		byr  string
		want bool
	}{
		{
			name: "lower bound valid iyr is valid",
			byr:  "1920",
			want: true,
		},
		{
			name: "upper bound valid iyr is valid",
			byr:  "2002",
			want: true,
		},
		{
			name: "midrange valid iyr is valid",
			byr:  "1960",
			want: true,
		},
		{
			name: "under lower bound invalid iyr is invalid",
			byr:  "1919",
			want: false,
		},
		{
			name: "above upper bound invalid iyr is invalid",
			byr:  "2003",
			want: false,
		},
		{
			name: "empty string iyr is invalid",
			byr:  "",
			want: false,
		},
		{
			name: "string iyr is invalid",
			byr:  "boo",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := identity{
				BYR: tt.byr,
			}

			assert.Equal(t, tt.want, i.byrValid())
		})
	}
}

func Test_identity_iyrValid(t *testing.T) {
	tests := []struct {
		name string
		iyr  string
		want bool
	}{
		{
			name: "lower bound valid iyr is valid",
			iyr:  "2010",
			want: true,
		},
		{
			name: "upper bound valid iyr is valid",
			iyr:  "2020",
			want: true,
		},
		{
			name: "midrange valid iyr is valid",
			iyr:  "2015",
			want: true,
		},
		{
			name: "under lower bound invalid iyr is invalid",
			iyr:  "2009",
			want: false,
		},
		{
			name: "above upper bound invalid iyr is invalid",
			iyr:  "2021",
			want: false,
		},
		{
			name: "empty string iyr is invalid",
			iyr:  "",
			want: false,
		},
		{
			name: "string iyr is invalid",
			iyr:  "boo",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := identity{
				IYR: tt.iyr,
			}

			assert.Equal(t, tt.want, i.iyrValid())
		})
	}
}

func Test_identity_hgtValid(t *testing.T) {
	tests := []struct {
		name string
		hgt  string
		want bool
	}{
		{
			name: "lower bound cm is valid",
			hgt:  "150cm",
			want: true,
		},
		{
			name: "upper bound cm is valid",
			hgt:  "193cm",
			want: true,
		},
		{
			name: "under lower bound cm is invalid",
			hgt:  "149cm",
			want: false,
		},
		{
			name: "above upper bound cm is invalid",
			hgt:  "194cm",
			want: false,
		},
		{
			name: "lower bound in is valid",
			hgt:  "59in",
			want: true,
		},
		{
			name: "upper bound in is valid",
			hgt:  "76in",
			want: true,
		},
		{
			name: "under lower bound in is invalid",
			hgt:  "58in",
			want: false,
		},
		{
			name: "above upper bound in is invalid",
			hgt:  "77in",
			want: false,
		},
		{
			name: "not a cm or in is invalid",
			hgt:  "boo",
			want: false,
		},
		{
			name: "empty hgt is invalid",
			hgt:  "",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := identity{
				HGT: tt.hgt,
			}
			assert.Equal(t, tt.want, i.hgtValid())
		})
	}
}
