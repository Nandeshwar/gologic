package main

func main() {

}

// * Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {

	/*
				// algo1: left - low value, right = high value .\
		       // check each node < right && node > left -- good condition otherwise return false

			    // algo2: inorder traversal- must be in ascending order otherwise return false
				 // algo 1
				    if root == nil {
				        return true
				    }

				    left := math.MinInt
				    right := math.MaxInt
				    return isValid(root, left, right)
	*/
	// algo 2
	if root == nil {
		return true
	}
	pVal := root.Val
	assigned := false
	return isValid2(root, &pVal, &assigned)
}

func isValid(root *TreeNode, left, right int) bool {
	if root == nil {
		return true
	}

	if !(root.Val < right && root.Val > left) {
		return false
	}

	return isValid(root.Left, left, root.Val) && isValid(root.Right, root.Val, right)
}

func isValid2(root *TreeNode, pVal *int, assigned *bool) bool {
	if root == nil {
		return true
	}

	left := isValid2(root.Left, pVal, assigned)

	if *assigned && root.Val <= *pVal {
		return false
	}
	*pVal = root.Val
	*assigned = true

	right := isValid2(root.Right, pVal, assigned)
	return left && right
}
