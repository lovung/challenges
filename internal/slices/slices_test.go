package slices

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		in  []any
		out []any
	}{
		{
			in:  []any{1, 2, 3},
			out: []any{3, 2, 1},
		},
		{
			in:  []any{"a", "b", "c"},
			out: []any{"c", "b", "a"},
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			Reverse(tt.in)
			assert.Equal(t, tt.out, tt.in)
		})
	}
}

func TestMaxValueAndIndex(t *testing.T) {
	type args struct {
		s []int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{
			name: "test1",
			args: args{
				s: []int{1, 2, 3, 4, 5},
			},
			want:  5,
			want1: 4,
		},
		{
			name: "test2",
			args: args{
				s: []int{10, 11, 100, -1, 5, 6},
			},
			want:  100,
			want1: 2,
		},
		{
			name: "empty",
			args: args{
				s: []int{},
			},
			want:  0,
			want1: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := MaxValueAndIndex(tt.args.s)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MaxValueAndIndex() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("MaxValueAndIndex() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
