package _17MergeTwoBinaryTrees

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
}

func (root *TreeNode)addLeftNode(node *TreeNode) {
	root.Left = node
}
func (root *TreeNode)addRightNode(node *TreeNode) {
	root.Right = node
}
func (root *TreeNode)addValue(v int) {
	root.Val = v
}

func (root *TreeNode)Print() {
	fmt.Println(root.Val)
	list := []*TreeNode{}
	if root.Left != nil {
		//root.Left.Print()
		list = append(list, root.Left)
	}
	if root.Right != nil {
		//root.Right.Print()
		list = append(list, root.Right)
	}
	for _, n := range list {
		n.Print()
	}
}
func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	rTree := &TreeNode{0, nil, nil}

	readTree(t1, rTree)
	readTree(t2, rTree)
	return rTree
}

func readTree(sourceNode *TreeNode, ret *TreeNode) {
	ret.Val += sourceNode.Val
	if sourceNode.Left != nil {
		if ret.Left == nil {
			ret.Left = &TreeNode{0 ,nil, nil}
		}
		readTree(sourceNode.Left, ret.Left)
	}
	if sourceNode.Right != nil {
		if ret.Right == nil {
			ret.Right = &TreeNode{0 ,nil, nil}
		}
		readTree(sourceNode.Right, ret.Right)
	}
}