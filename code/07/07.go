package main

func main() {

}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	var left, right []int
	for k, _ := range inorder {
		if preorder[0] == inorder[k] {
			left = inorder[0:k]
			right = inorder[k+1 : len(inorder)]
			break
		}
	}
	return &TreeNode{
		Val:   preorder[0],
		Left:  buildTree(preorder[1:len(left)+1], left),
		Right: buildTree(preorder[len(left)+1:], right),
	}
}
