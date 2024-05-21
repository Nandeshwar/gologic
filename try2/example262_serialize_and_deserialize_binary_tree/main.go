package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	var strArr []string
	serializeHelper(root, &strArr)
	fmt.Println(strArr)
	return strings.Join(strArr, ",")
}

// preorder traversal and create string
func serializeHelper(root *TreeNode, strArr *[]string) {
	if root == nil {
		*strArr = append(*strArr, "n")
		return
	}
	*strArr = append(*strArr, strconv.Itoa(root.Val))
	serializeHelper(root.Left, strArr)
	serializeHelper(root.Right, strArr)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {

	ind := -1
	return deserializeHelper(strings.Split(data, ","), &ind)
}

// depth first search and create node
func deserializeHelper(data []string, ind *int) *TreeNode {
	*ind = *ind + 1

	if string(data[*ind]) == "n" {
		return nil
	}

	d, _ := strconv.Atoi(string(data[*ind]))
	node := &TreeNode{
		Val:   d,
		Left:  deserializeHelper(data, ind),
		Right: deserializeHelper(data, ind),
	}
	return node
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
