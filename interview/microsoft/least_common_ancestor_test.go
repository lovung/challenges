package microsoft

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_LeastCommonAncestor(t *testing.T) {
	t.Run("all", func(t *testing.T) {
		assert.Nil(t, LeastCommonAncestor(nil, nil, nil))
		assert.Nil(t, LeastCommonAncestor2(nil, nil, nil))
		// Wrong case: assert.Nil(t, LeastCommonAncestor(&n1, &n1, &n3))
	})
	t.Run("sol_1_right", func(t *testing.T) {
		n4 := TreeNode{Val: 4}
		n5 := TreeNode{Val: 5}
		n6 := TreeNode{Val: 6}
		n7 := TreeNode{Val: 7}
		n2 := TreeNode{Val: 2, Left: &n4, Right: &n5}
		n3 := TreeNode{Val: 3, Left: &n7, Right: &n6}
		n1 := TreeNode{Val: 1, Left: &n2, Right: &n3}
		assert.Equal(t, 2, LeastCommonAncestor(&n1, &n4, &n5).Val)
		assert.Equal(t, 1, LeastCommonAncestor(&n1, &n4, &n2).Val)
		assert.Equal(t, 1, LeastCommonAncestor(&n1, &n4, &n6).Val)
		assert.Equal(t, 1, LeastCommonAncestor(&n1, &n4, &n3).Val)
		assert.Equal(t, 1, LeastCommonAncestor(&n1, &n6, &n5).Val)
		assert.Equal(t, 3, LeastCommonAncestor(&n1, &n6, &n7).Val)
		assert.Equal(t, 1, LeastCommonAncestor(&n1, &n6, &n3).Val)
		// Wrong case: assert.Nil(t, LeastCommonAncestor(&n1, &n1, &n3))
	})
	t.Run("sol_2", func(t *testing.T) {
		n4 := TreeNode{Val: 4}
		n5 := TreeNode{Val: 5}
		n6 := TreeNode{Val: 6}
		n7 := TreeNode{Val: 7}
		n2 := TreeNode{Val: 2, Left: &n4, Right: &n5}
		n3 := TreeNode{Val: 3, Left: &n7, Right: &n6}
		n1 := TreeNode{Val: 1, Left: &n2, Right: &n3}
		assert.Equal(t, 2, LeastCommonAncestor2(&n1, &n4, &n5).Val)
		assert.Equal(t, 1, LeastCommonAncestor2(&n1, &n4, &n2).Val)
		assert.Equal(t, 1, LeastCommonAncestor2(&n1, &n4, &n6).Val)
		assert.Equal(t, 1, LeastCommonAncestor2(&n1, &n4, &n3).Val)
		assert.Equal(t, 1, LeastCommonAncestor2(&n1, &n6, &n5).Val)
		assert.Equal(t, 3, LeastCommonAncestor(&n1, &n6, &n7).Val)
		assert.Equal(t, 1, LeastCommonAncestor2(&n1, &n6, &n3).Val)
		assert.Nil(t, LeastCommonAncestor2(&n1, &n1, &n3))
	})

}
