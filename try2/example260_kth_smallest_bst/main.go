package main

func main() {

}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {

	counter := 0
	return findKthSmallest(root, k, &counter)

}

func findKthSmallest(root *TreeNode, k int, counter *int) int {
	if root == nil {
		return 0
	}

	left := findKthSmallest(root.Left, k, counter)
	*counter = *counter + 1
	if *counter == k {
		return root.Val
	}

	right := findKthSmallest(root.Right, k, counter)
	return left + right

}
