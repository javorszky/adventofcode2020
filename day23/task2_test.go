package day23

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_generateTenPlusNCups(t *testing.T) {
	type args struct {
		in []int
		n  int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "generates a 1 extra long cup circle",
			args: args{
				in: []int{5, 6, 3, 2, 1, 4},
				n:  1,
			},
			want: []int{5, 6, 3, 2, 1, 4, 7},
		},
		{
			name: "generates a 10 extra long cup circle",
			args: args{
				in: []int{5, 6, 3, 2, 1, 4},
				n:  10,
			},
			want: []int{5, 6, 3, 2, 1, 4, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, generateTenPlusNCups(tt.args.in, tt.args.n))
		})
	}
}

func Test_productOfCupsClockwiseToOne(t *testing.T) {
	type args struct {
		in []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "returns product for slice where 1 is at the beginning",
			args: args{
				in: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0},
			},
			want: 6,
		},
		{
			name: "returns product for slice where 1 is somewhere in the middle",
			args: args{
				in: []int{6, 7, 1, 8, 9, 0, 2, 3, 4, 5},
			},
			want: 72,
		},
		{
			name: "returns product for slice where 1 is somewhere in the middle, and one of the numbers is 0",
			args: args{
				in: []int{6, 7, 8, 1, 9, 0, 2, 3, 4, 5},
			},
			want: 0,
		},
		{
			name: "returns product for slice where 1 is the last but one",
			args: args{
				in: []int{6, 7, 8, 9, 0, 2, 3, 4, 1, 5},
			},
			want: 30,
		},
		{
			name: "returns product for slice where 1 is the last but one",
			args: args{
				in: []int{6, 7, 8, 9, 0, 2, 3, 4, 1},
			},
			want: 42,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, productOfCupsClockwiseToOne(tt.args.in))
		})
	}
}

func Test_newBigOof(t *testing.T) {
	type args struct {
		input []int
		cups  int
	}
	tests := []struct {
		name string
		args args
		want bigOof
	}{
		{
			name: "generates input based on file and number of cups total",
			args: args{
				input: []int{1, 3, 5, 4, 6, 8, 7, 2, 9},
				cups:  12,
			},
			want: bigOof{
				length:  12,
				current: 1,
				whatsOn: map[int]int{
					1:  1,
					2:  3,
					3:  5,
					4:  4,
					5:  6,
					6:  8,
					7:  7,
					8:  2,
					9:  9,
					10: 10,
					11: 11,
					12: 12,
				},
				whereIs: map[int]int{
					1:  1,
					2:  8,
					3:  2,
					4:  4,
					5:  3,
					6:  5,
					7:  7,
					8:  6,
					9:  9,
					10: 10,
					11: 11,
					12: 12,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, newBigOof(tt.args.input, tt.args.cups))
		})
	}
}

func Test_bigOof_step(t *testing.T) {
	type fields struct {
		whatsOn map[int]int
		whereIs map[int]int
		current int
		length  int
	}
	tests := []struct {
		name   string
		fields fields
		want   bigOof
	}{
		//{
		//	name: "correctly calculates step when next number is lower",
		//	fields: fields{
		//		whatsOn: map[int]int{
		//			1: 6,
		//			2: 8,
		//			3: 4,
		//			4: 7,
		//			5: 3,
		//			6: 2,
		//			7: 1,
		//			8: 9,
		//			9: 5,
		//		},
		//		whereIs: map[int]int{
		//			1: 7,
		//			2: 6,
		//			3: 5,
		//			4: 3,
		//			5: 9,
		//			6: 1,
		//			7: 4,
		//			8: 2,
		//			9: 8,
		//		},
		//		current: 3,
		//		length:  9,
		//	},
		//	want: bigOof{
		//		whatsOn: map[int]int{
		//			1: 6,
		//			2: 8,
		//			3: 2,
		//			4: 1,
		//			5: 9,
		//			6: 4,
		//			7: 7,
		//			8: 3,
		//			9: 5,
		//		},
		//		whereIs: map[int]int{
		//			1: 4,
		//			2: 3,
		//			3: 8,
		//			4: 6,
		//			5: 9,
		//			6: 1,
		//			7: 7,
		//			8: 2,
		//			9: 5,
		//		},
		//		current: 5,
		//		length:  9,
		//	},
		//},
		//{
		//	name: "correctly calculates step when next number is lower and pickup wraps",
		//	fields: fields{
		//		whatsOn: map[int]int{
		//			1: 9,
		//			2: 5,
		//			3: 6,
		//			4: 8,
		//			5: 4,
		//			6: 7,
		//			7: 3,
		//			8: 2,
		//			9: 1,
		//		},
		//		whereIs: map[int]int{
		//			1: 9,
		//			2: 8,
		//			3: 7,
		//			4: 5,
		//			5: 2,
		//			6: 3,
		//			7: 6,
		//			8: 4,
		//			9: 1,
		//		},
		//		current: 3,
		//		length:  9,
		//	},
		//	want: bigOof{
		//		whatsOn: map[int]int{
		//			1: 3,
		//			2: 5,
		//			3: 6,
		//			4: 8,
		//			5: 2,
		//			6: 1,
		//			7: 9,
		//			8: 4,
		//			9: 7,
		//		},
		//		whereIs: map[int]int{
		//			1: 6,
		//			2: 5,
		//			3: 1,
		//			4: 8,
		//			5: 2,
		//			6: 3,
		//			7: 9,
		//			8: 4,
		//			9: 7,
		//		},
		//		current: 5,
		//		length:  9,
		//	},
		//},
		{
			name: "correctly calculates step when next number is higher",
			fields: fields{
				whatsOn: map[int]int{
					1: 1,
					2: 3,
					3: 8,
					4: 6,
					5: 4,
					6: 5,
					7: 7,
					8: 2,
					9: 9,
				},
				whereIs: map[int]int{
					1: 1,
					2: 8,
					3: 2,
					4: 5,
					5: 6,
					6: 5,
					7: 7,
					8: 3,
					9: 9,
				},
				current: 3,
				length:  9,
			},
			want: bigOof{
				whatsOn: map[int]int{
					1: 1,
					2: 3,
					3: 5,
					4: 7,
					5: 2,
					6: 8,
					7: 6,
					8: 4,
					9: 9,
				},
				whereIs: map[int]int{
					1: 1,
					2: 5,
					3: 2,
					4: 8,
					5: 3,
					6: 7,
					7: 4,
					8: 6,
					9: 9,
				},
				current: 5,
				length:  9,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := bigOof{
				whatsOn: tt.fields.whatsOn,
				whereIs: tt.fields.whereIs,
				current: tt.fields.current,
				length:  tt.fields.length,
			}
			assert.Equal(t, tt.want, b.step())
		})
	}
}
