package day25

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findLoopSize(t *testing.T) {
	type args struct {
		what    int
		subject int
		target  int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "finds loop size for example in day 25 ",
			args: args{
				what:    1,
				subject: 7,
				target:  5764801,
			},
			want:    8,
			wantErr: false,
		},
		{
			name: "finds loop size for other example in day 25 ",
			args: args{
				what:    1,
				subject: 7,
				target:  17807724,
			},
			want:    11,
			wantErr: false,
		},
		{
			name: "can't find loop size for number",
			args: args{
				what:    1,
				subject: 7,
				target:  17807723,
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := findLoopSize(tt.args.what, tt.args.subject, tt.args.target)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tt.want, got)
		})
	}
}
