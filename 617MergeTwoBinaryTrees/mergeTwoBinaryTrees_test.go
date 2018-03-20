package _17MergeTwoBinaryTrees

import (
	"testing"
)

func Test_MergeTwoBinaryTrees(t *testing.T) {
	tree1 := &TreeNode{1, nil, nil}
	tree1.Left = &TreeNode{3, nil, nil}
	tree1.Right = &TreeNode{2, nil, nil}
	tree1.Left.addLeftNode(&TreeNode{5, nil, nil})

	tree2 := &TreeNode{2, nil, nil}
	tree2.Left = &TreeNode{1, nil, nil}
	tree2.Left.Right = &TreeNode{4, nil, nil}
	tree2.Right = &TreeNode{3, nil, &TreeNode{7, nil, nil}}

	r := mergeTrees(tree1, tree2)
	r.Print()
}
