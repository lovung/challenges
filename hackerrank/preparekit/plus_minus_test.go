package preparekit

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_plusMinus(t *testing.T) {
	type args struct {
		arr []int32
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				arr: []int32{1, 1, 0, -1, -1},
			},
			want: `0.400000
0.400000
0.200000
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var output bytes.Buffer
			w = &output
			plusMinus(tt.args.arr)
			assert.Equal(t, tt.want, output.String())
		})
	}
}
