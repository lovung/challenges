package letters

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_partition(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		// {
		// 	args: args{
		// 		s: "aab",
		// 	},
		// 	want: [][]string{
		// 		{"a", "a", "b"},
		// 		{"aa", "b"},
		// 	},
		// },
		{
			args: args{
				s: "abaeeada",
			},
			want: [][]string{
				{"a", "b", "a", "e", "e", "a", "d", "a"},
				{"a", "b", "a", "e", "e", "ada"},
				{"a", "b", "a", "ee", "a", "d", "a"},
				{"a", "b", "a", "ee", "ada"},
				{"a", "b", "aeea", "d", "a"},
				{"aba", "e", "e", "a", "d", "a"},
				{"aba", "e", "e", "ada"},
				{"aba", "ee", "a", "d", "a"},
				{"aba", "ee", "ada"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, partition(tt.args.s))
		})
	}
}
