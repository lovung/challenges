package binarytree

import (
	"reflect"
	"testing"

	"github.com/lovung/ds/pointer"
	"github.com/lovung/ds/trees"
)

func Test_averageOfLevels(t *testing.T) {
	type args struct {
		root *trees.TreeNode[int]
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		{
			name: "test1",
			args: args{
				root: trees.Slice2TreeNode([]*int{
					pointer.Of(3),
					pointer.Of(9), pointer.Of(20),
					nil, nil, pointer.Of(15), pointer.Of(7),
				}),
			},
			want: []float64{3, 14.5, 11},
		},
		{
			name: "test2",
			args: args{
				root: trees.Slice2TreeNode([]*int{
					pointer.Of(3),
					pointer.Of(9), pointer.Of(20),
					pointer.Of(15), pointer.Of(7),
				}),
			},
			want: []float64{3, 14.5, 11},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := averageOfLevels(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("averageOfLevels() = %v, want %v", got, tt.want)
			}
		})
	}
}
