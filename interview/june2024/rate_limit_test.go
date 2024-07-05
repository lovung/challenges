package june2024

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getRequestStatus(t *testing.T) {
	type args struct {
		requests []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			args: args{
				requests: []string{
					"www.abc.com",
					"www.hd.com",
					"www.abc.com",
					"www.pqr.com",
					"www.abc.com",
					"www.pqr.com",
					"www.pqr.com",
					"www.abc.com",
				},
			},
			want: []string{
				"{status: 200, message: OK}",
				"{status: 200, message: OK}",
				"{status: 200, message: OK}",
				"{status: 200, message: OK}",
				"{status: 429, message: Too many requests}",
				"{status: 200, message: OK}",
				"{status: 429, message: Too many requests}",
				"{status: 200, message: OK}",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getRequestStatus(tt.args.requests)
			assert.Equal(t, tt.want, got)
		})
	}
}
