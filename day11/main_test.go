package day11

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_t1normalize(t *testing.T) {
	// input grid that is the basis of this is
	//    0 1 2 3 4 5
	//  0 1 2 3 4 5 6
	//  6 7 8 9 0 a b
	// 12 c d e f g h
	// 18 i j k l m n
	const (
		lineLength = 6
		input      = "1234567890abcdefghijklmn"
	)
	type args struct {
		input []string
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 string
	}{
		{
			name: "normalizes grid into a line length and a string",
			args: args{
				input: []string{
					"123456",
					"7890ab",
					"cdefgh",
					"ijklmn",
				},
			},
			want:  lineLength,
			want1: input,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := t1normalize(tt.args.input)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want1, got1)
		})
	}
}

func Test_tl(t *testing.T) {
	// input grid that is the basis of this is
	//    0 1 2 3 4 5
	//  0 1 2 3 4 5 6
	//  6 7 8 9 0 a b
	// 12 c d e f g h
	// 18 i j k l m n
	const (
		lineLength = 6
		input      = "1234567890abcdefghijklmn"
	)
	type args struct {
		c          int
		lineLength int
		input      string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "returns empty top left for character in top row left column",
			args: args{
				c:          0,
				lineLength: lineLength,
				input:      input,
			},
			want: "",
		},
		{
			name: "returns empty top left for character in top row right column",
			args: args{
				c:          5,
				lineLength: lineLength,
				input:      input,
			},
			want: "",
		},
		{
			name: "returns empty top left for character in second row left column",
			args: args{
				c:          6,
				lineLength: lineLength,
				input:      input,
			},
			want: "",
		},
		{
			name: "returns correct neighbour top left for character in second row right column",
			args: args{
				c:          11,
				lineLength: lineLength,
				input:      input,
			},
			want: "5",
		},
		{
			name: "returns correct neighbour top left for character in bottom row right column",
			args: args{
				c:          23,
				lineLength: lineLength,
				input:      input,
			},
			want: "g",
		},
		{
			name: "returns empty neighbour top left for character in bottom row left column",
			args: args{
				c:          18,
				lineLength: lineLength,
				input:      input,
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, c := topLeft(tt.args.c, tt.args.lineLength, tt.args.input)
			assert.Equal(t, tt.want, c)
		})
	}
}

func Test_top(t *testing.T) {
	// input grid that is the basis of this is
	//    0 1 2 3 4 5
	//  0 1 2 3 4 5 6
	//  6 7 8 9 0 a b
	// 12 c d e f g h
	// 18 i j k l m n
	const (
		lineLength = 6
		input      = "1234567890abcdefghijklmn"
	)
	type args struct {
		c          int
		lineLength int
		input      string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "returns empty top for character in top row left column",
			args: args{
				c:          0,
				lineLength: lineLength,
				input:      input,
			},
			want: "",
		},
		{
			name: "returns empty top for character in top row right column",
			args: args{
				c:          5,
				lineLength: lineLength,
				input:      input,
			},
			want: "",
		},
		{
			name: "returns correct neighbour top for character in second row left column",
			args: args{
				c:          6,
				lineLength: lineLength,
				input:      input,
			},
			want: "1",
		},
		{
			name: "returns correct neighbour top for character in second row right column",
			args: args{
				c:          11,
				lineLength: lineLength,
				input:      input,
			},
			want: "6",
		},
		{
			name: "returns correct neighbour top for character in bottom row right column",
			args: args{
				c:          23,
				lineLength: lineLength,
				input:      input,
			},
			want: "h",
		},
		{
			name: "returns empty neighbour top for character in bottom row left column",
			args: args{
				c:          18,
				lineLength: lineLength,
				input:      input,
			},
			want: "c",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, c := top(tt.args.c, tt.args.lineLength, tt.args.input)
			assert.Equal(t, tt.want, c)
		})
	}
}

func Test_topRight(t *testing.T) {
	/// input grid that is the basis of this is
	//    0 1 2 3 4 5
	//  0 1 2 3 4 5 6
	//  6 7 8 9 0 a b
	// 12 c d e f g h
	// 18 i j k l m n
	const (
		lineLength = 6
		input      = "1234567890abcdefghijklmn"
	)
	type args struct {
		c          int
		lineLength int
		input      string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "returns empty top right for character in top row left column",
			args: args{
				c:          0,
				lineLength: lineLength,
				input:      input,
			},
			want: "",
		},
		{
			name: "returns empty top right for character in top row right column",
			args: args{
				c:          5,
				lineLength: lineLength,
				input:      input,
			},
			want: "",
		},
		{
			name: "returns correct neighbour top right for character in second row left column",
			args: args{
				c:          6,
				lineLength: lineLength,
				input:      input,
			},
			want: "2",
		},
		{
			name: "returns empty top right for character in second row right column",
			args: args{
				c:          11,
				lineLength: lineLength,
				input:      input,
			},
			want: "",
		},
		{
			name: "returns empty top right for character in bottom row right column",
			args: args{
				c:          23,
				lineLength: lineLength,
				input:      input,
			},
			want: "",
		},
		{
			name: "returns correct neighbour top right for character in bottom row left column",
			args: args{
				c:          18,
				lineLength: lineLength,
				input:      input,
			},
			want: "d",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, c := topRight(tt.args.c, tt.args.lineLength, tt.args.input)
			assert.Equal(t, tt.want, c)
		})
	}
}

func Test_right(t *testing.T) {
	// input grid that is the basis of this is
	//      0 1 2 3 4 5
	//    o------------
	//  0 | 1 2 3 4 5 6
	//  6 | 7 8 9 0 a b
	// 12 | c d e f g h
	// 18 | i j k l m n
	const (
		lineLength = 6
		input      = "1234567890abcdefghijklmn"
	)
	type args struct {
		c          int
		lineLength int
		input      string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "returns correct neighbour right for character in top row left column",
			args: args{
				c:          0,
				lineLength: lineLength,
				input:      input,
			},
			want: "2",
		},
		{
			name: "returns empty right for character in top row right column",
			args: args{
				c:          5,
				lineLength: lineLength,
				input:      input,
			},
			want: "",
		},
		{
			name: "returns correct neighbour right for character in second row left column",
			args: args{
				c:          6,
				lineLength: lineLength,
				input:      input,
			},
			want: "8",
		},
		{
			name: "returns empty right for character in second row right column",
			args: args{
				c:          11,
				lineLength: lineLength,
				input:      input,
			},
			want: "",
		},
		{
			name: "returns empty right for character in bottom row right column",
			args: args{
				c:          23,
				lineLength: lineLength,
				input:      input,
			},
			want: "",
		},
		{
			name: "returns correct neighbour right for character in bottom row left column",
			args: args{
				c:          18,
				lineLength: lineLength,
				input:      input,
			},
			want: "j",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, c := right(tt.args.c, tt.args.lineLength, tt.args.input)
			assert.Equal(t, tt.want, c)
		})
	}
}

func Test_bottomRight(t *testing.T) {
	// input grid that is the basis of this is
	//      0 1 2 3 4 5
	//    o------------
	//  0 | 1 2 3 4 5 6
	//  6 | 7 8 9 0 a b
	// 12 | c d e f g h
	// 18 | i j k l m n
	const (
		lineLength = 6
		input      = "1234567890abcdefghijklmn"
	)
	type args struct {
		c          int
		lineLength int
		input      string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "returns correct bottom right for character in top row left column",
			args: args{
				c:          0,
				lineLength: lineLength,
				input:      input,
			},
			want: "8",
		},
		{
			name: "returns empty bottom right for character in top row right column",
			args: args{
				c:          5,
				lineLength: lineLength,
				input:      input,
			},
			want: "",
		},
		{
			name: "returns correct neighbour bottom right for character in second row left column",
			args: args{
				c:          6,
				lineLength: lineLength,
				input:      input,
			},
			want: "d",
		},
		{
			name: "returns empty bottom right for character in second row right column",
			args: args{
				c:          11,
				lineLength: lineLength,
				input:      input,
			},
			want: "",
		},
		{
			name: "returns empty bottom right for character in bottom row right column",
			args: args{
				c:          23,
				lineLength: lineLength,
				input:      input,
			},
			want: "",
		},
		{
			name: "returns empty bottom right for character in bottom row left column",
			args: args{
				c:          18,
				lineLength: lineLength,
				input:      input,
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, c := bottomRight(tt.args.c, tt.args.lineLength, tt.args.input)
			assert.Equal(t, tt.want, c)
		})
	}
}

func Test_bottom(t *testing.T) {
	// input grid that is the basis of this is
	//      0 1 2 3 4 5
	//    o------------
	//  0 | 1 2 3 4 5 6
	//  6 | 7 8 9 0 a b
	// 12 | c d e f g h
	// 18 | i j k l m n
	const (
		lineLength = 6
		input      = "1234567890abcdefghijklmn"
	)
	type args struct {
		c          int
		lineLength int
		input      string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "returns correct bottom for character in top row left column",
			args: args{
				c:          0,
				lineLength: lineLength,
				input:      input,
			},
			want: "7",
		},
		{
			name: "returns correct bottom for character in top row right column",
			args: args{
				c:          5,
				lineLength: lineLength,
				input:      input,
			},
			want: "b",
		},
		{
			name: "returns correct bottom for character in second row left column",
			args: args{
				c:          6,
				lineLength: lineLength,
				input:      input,
			},
			want: "c",
		},
		{
			name: "returns correct bottom for character in second row right column",
			args: args{
				c:          11,
				lineLength: lineLength,
				input:      input,
			},
			want: "h",
		},
		{
			name: "returns empty bottom for character in bottom row right column",
			args: args{
				c:          23,
				lineLength: lineLength,
				input:      input,
			},
			want: "",
		},
		{
			name: "returns empty bottom for character in bottom row left column",
			args: args{
				c:          18,
				lineLength: lineLength,
				input:      input,
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, c := bottom(tt.args.c, tt.args.lineLength, tt.args.input)
			assert.Equal(t, tt.want, c)
		})
	}
}

func Test_bottomLeft(t *testing.T) {
	// input grid that is the basis of this is
	//      0 1 2 3 4 5
	//    o------------
	//  0 | 1 2 3 4 5 6
	//  6 | 7 8 9 0 a b
	// 12 | c d e f g h
	// 18 | i j k l m n
	const (
		lineLength = 6
		input      = "1234567890abcdefghijklmn"
	)
	type args struct {
		c          int
		lineLength int
		input      string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "returns empty bottom left for character in top row left column",
			args: args{
				c:          0,
				lineLength: lineLength,
				input:      input,
			},
			want: "",
		},
		{
			name: "returns correct bottom left for character in top row right column",
			args: args{
				c:          5,
				lineLength: lineLength,
				input:      input,
			},
			want: "a",
		},
		{
			name: "returns empty bottom left for character in second row left column",
			args: args{
				c:          6,
				lineLength: lineLength,
				input:      input,
			},
			want: "",
		},
		{
			name: "returns correct bottom left for character in second row right column",
			args: args{
				c:          11,
				lineLength: lineLength,
				input:      input,
			},
			want: "g",
		},
		{
			name: "returns empty bottom left for character in bottom row right column",
			args: args{
				c:          23,
				lineLength: lineLength,
				input:      input,
			},
			want: "",
		},
		{
			name: "returns empty bottom left for character in bottom row left column",
			args: args{
				c:          18,
				lineLength: lineLength,
				input:      input,
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, c := bottomLeft(tt.args.c, tt.args.lineLength, tt.args.input)
			assert.Equal(t, tt.want, c)
		})
	}
}

func Test_left(t *testing.T) {
	// input grid that is the basis of this is
	//      0 1 2 3 4 5
	//    o------------
	//  0 | 1 2 3 4 5 6
	//  6 | 7 8 9 0 a b
	// 12 | c d e f g h
	// 18 | i j k l m n
	const (
		lineLength = 6
		input      = "1234567890abcdefghijklmn"
	)
	type args struct {
		c          int
		lineLength int
		input      string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "returns empty left for character in top row left column",
			args: args{
				c:          0,
				lineLength: lineLength,
				input:      input,
			},
			want: "",
		},
		{
			name: "returns correct left for character in top row right column",
			args: args{
				c:          5,
				lineLength: lineLength,
				input:      input,
			},
			want: "5",
		},
		{
			name: "returns empty left for character in second row left column",
			args: args{
				c:          6,
				lineLength: lineLength,
				input:      input,
			},
			want: "",
		},
		{
			name: "returns correct left for character in second row right column",
			args: args{
				c:          11,
				lineLength: lineLength,
				input:      input,
			},
			want: "a",
		},
		{
			name: "returns correct left for character in bottom row right column",
			args: args{
				c:          23,
				lineLength: lineLength,
				input:      input,
			},
			want: "m",
		},
		{
			name: "returns empty left for character in bottom row left column",
			args: args{
				c:          18,
				lineLength: lineLength,
				input:      input,
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, c := left(tt.args.c, tt.args.lineLength, tt.args.input)
			assert.Equal(t, tt.want, c)
		})
	}
}
