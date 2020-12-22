package day22

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_slicePop(t *testing.T) {
	type args struct {
		in []int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		want1   []int
		wantErr bool
	}{
		{
			name: "returns error if nil slice is passed in",
			args: args{
				in: nil,
			},
			want:    0,
			want1:   nil,
			wantErr: true,
		},
		{
			name: "returns error if empty slice is passed in",
			args: args{
				in: []int{},
			},
			want:    0,
			want1:   nil,
			wantErr: true,
		},
		{
			name: "returns first element of a one element slice, and an empty slice",
			args: args{
				in: []int{1},
			},
			want:    1,
			want1:   []int{},
			wantErr: false,
		},
		{
			name: "returns first element of a two element slice, and the rest of the slice",
			args: args{
				in: []int{1, 2, 3, 4},
			},
			want:    1,
			want1:   []int{2, 3, 4},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := slicePop(tt.args.in)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want1, got1)
		})
	}
}
