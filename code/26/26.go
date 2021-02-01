package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 树的子结构
// 递归解法
func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}
	if A == nil && B == nil {
		return true
	}
	return tree(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}

func tree(left, right *TreeNode) bool {
	if right == nil {
		return true
	}
	if left == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}
	return tree(left.Left, right.Left) && tree(left.Right, right.Right)
}
