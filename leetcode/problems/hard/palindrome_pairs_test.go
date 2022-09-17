package hard

import (
	"reflect"
	"testing"
)

func Test_palindromePairs(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			args: args{
				words: []string{"abcd", "dcba", "lls", "s", "sssll"},
			},
			want: [][]int{{0, 1}, {1, 0}, {2, 4}, {3, 2}},
		},
		{
			args: args{
				words: []string{"abcd", "dcba", "lls", "s", "sssll", "abcd"},
			},
			want: [][]int{{0, 1}, {1, 0}, {1, 5}, {2, 4}, {3, 2}, {5, 1}},
		},
		{
			args: args{
				words: []string{"bat", "tab", "cat"},
			},
			want: [][]int{{0, 1}, {1, 0}},
		},
		{
			args: args{
				words: []string{"a", ""},
			},
			want: [][]int{{0, 1}, {1, 0}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := palindromePairs(tt.args.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("palindromePairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
