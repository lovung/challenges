package jul2024

import (
	"testing"

	"github.com/lovung/ds/pointer"
	"github.com/lovung/ds/trees"
)

func Test_countPairs(t *testing.T) {
	type args struct {
		root     *trees.TreeNode[int]
		distance int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				root: trees.Slice2TreeNode([]*int{
					pointer.Of(1), pointer.Of(2), pointer.Of(3), nil, pointer.Of(4),
				}),
				distance: 3,
			},
			want: 1,
		},
		{
			args: args{
				root: trees.Slice2TreeNode([]*int{
					pointer.Of(1),
					pointer.Of(2),
					pointer.Of(3),
					pointer.Of(4),
					pointer.Of(5),
					pointer.Of(6),
					pointer.Of(7),
				}),
				distance: 3,
			},
			want: 2,
		},
		{
			// 7,1,4,6,null,5,3,null,null,null,null,null,2
			args: args{
				root: trees.Slice2TreeNode([]*int{
					pointer.Of(7),
					pointer.Of(1),
					pointer.Of(4),
					pointer.Of(6),
					nil,
					pointer.Of(5),
					pointer.Of(3),
					nil, nil, nil, nil, nil,
					pointer.Of(2),
				}),
				distance: 3,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countPairs(tt.args.root, tt.args.distance); got != tt.want {
				t.Errorf("countPairs() = %v, want %v", got, tt.want)
			}
		})
		t.Run(tt.name, func(t *testing.T) {
			if got := countPairs2(tt.args.root, tt.args.distance); got != tt.want {
				t.Errorf("countPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
