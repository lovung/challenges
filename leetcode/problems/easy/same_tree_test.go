package easy

import (
	"testing"

	"github.com/lovung/challenges/internal/pointer"
	"github.com/lovung/challenges/internal/tree"
	"github.com/stretchr/testify/assert"
)

func Test_isSameTree(t *testing.T) {
	type args struct {
		p *tree.TreeNode[int]
		q *tree.TreeNode[int]
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "empty-both",
			args: args{
				p: nil,
				q: nil,
			},
			want: true,
		},
		{
			name: "empty-1",
			args: args{
				p: nil,
				q: tree.Slice2TreeNode([]*int{pointer.Of(1)}),
			},
			want: false,
		},
		{
			name: "diff-1",
			args: args{
				p: tree.Slice2TreeNode([]*int{pointer.Of(1)}),
				q: tree.Slice2TreeNode([]*int{pointer.Of(2)}),
			},
			want: false,
		},
		{
			name: "normal-equal",
			args: args{
				p: tree.Slice2TreeNode([]*int{pointer.Of(1), pointer.Of(2), pointer.Of(3)}),
				q: tree.Slice2TreeNode([]*int{pointer.Of(1), pointer.Of(2), pointer.Of(3)}),
			},
			want: true,
		},
		{
			name: "nil-left-and-nil-right",
			args: args{
				p: tree.Slice2TreeNode([]*int{pointer.Of(1), nil, pointer.Of(3)}),
				q: tree.Slice2TreeNode([]*int{pointer.Of(1), pointer.Of(3)}),
			},
			want: false,
		},
		{
			name: "1-nil-right",
			args: args{
				p: tree.Slice2TreeNode([]*int{pointer.Of(1), pointer.Of(2), pointer.Of(3)}),
				q: tree.Slice2TreeNode([]*int{pointer.Of(1), pointer.Of(2)}),
			},
			want: false,
		},
		{
			name: "diff-value",
			args: args{
				p: tree.Slice2TreeNode([]*int{pointer.Of(1), pointer.Of(1), pointer.Of(2)}),
				q: tree.Slice2TreeNode([]*int{pointer.Of(1), pointer.Of(2), pointer.Of(1)}),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, isSameTree(tt.args.p, tt.args.q))
		})
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, isSameTree2(tt.args.p, tt.args.q))
		})
	}
}
