package main

import "math"

func main() {

}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	result := math.MinInt
	maxPathSumHelper(root, &result)
	return result
}

func maxPathSumHelper(root *TreeNode, result *int) int {
	if root == nil {
		return 0
	}

	leftMax := Max(0, maxPathSumHelper(root.Left, result))
	rightMax := Max(0, maxPathSumHelper(root.Right, result))

	// triangle  root + left value + right value   : will be path
	*result = Max(*result, root.Val+leftMax+rightMax)

	// returns node val + max of left and right:
	return root.Val + Max(leftMax, rightMax)
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
