// https://www.youtube.com/watch?v=nBbYMdtZIuc&list=PLU_sdQYzUj2keVENTP0a5rdykRSgg9Wp-&index=16
package main
import "fmt"

type TreeNode12 struct {
	Item int
	Left *TreeNode12
	Right *TreeNode12
}



func main() {
	var expectedSum = 8
	node4 := TreeNode12{Item: 4}
	node5 := TreeNode12{Item: 5}
	node6 := TreeNode12{Item: 6}

	node2 := TreeNode12{Item: 2, Left: &node4, Right: &node5}

	node3 := TreeNode12{Item: 3, Left: &node6}

	node1 := TreeNode12{Item: 1, Left: &node2, Right: &node3}

	root := &node1
	treeSum(root, expectedSum)
}


func treeSum(root *TreeNode12, expectedSum int) {
	if root == nil {
		return 
	}
	expectedSum -= root.Item;
	if expectedSum == 0 {
		fmt.Println("Found")
	}
	treeSum(root.Left, expectedSum)
	treeSum(root.Right, expectedSum)
	return
}