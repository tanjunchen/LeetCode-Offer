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

// 对称的二叉树
// 递归解法
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return dfsIsSymmetric(root.Left, root.Right)
}

func dfsIsSymmetric(l, r *TreeNode) bool {
	if l == nil && r == nil {
		return true
	}
	if l == nil && r != nil || l != nil && r == nil || l.Val != r.Val {
		return false
	}
	return dfsIsSymmetric(l.Left, r.Right) && dfsIsSymmetric(l.Right, r.Left)
}

func isSymmetric2(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var queue []*TreeNode
	queue = append(queue, root)
	for len(queue) != 0 {
		for i := len(queue); i > 0; i-- {
			temp := queue[0]
			queue = queue[1:]
			if temp != nil {
				queue = append(queue, temp.Left)
				queue = append(queue, temp.Right)
			}
		}
		for i := 0; i < len(queue)/2; i++ {
			l, r := queue[i], queue[len(queue)-1-i]
			if l != nil && r != nil && l.Val != r.Val {
				return false
			}
			if l == nil && r != nil || l != nil && r == nil {
				return false
			}
		}
	}
	return true
}
