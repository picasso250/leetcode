package main

//   Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	left := treeToList(root.Left)
	for _, x := range left {
		if !(x < root.Val) {
			return false
		}
	}
	right := treeToList(root.Right)
	for _, x := range right {
		if !(x > root.Val) {
			return false
		}
	}
	return isValidBST(root.Left) && isValidBST(root.Right)
}
func treeToList(root *TreeNode) (ret []int) {
	if root == nil {
		return
	}
	ret = []int{root.Val}
	ret = append(ret, treeToList(root.Left)...)
	ret = append(ret, treeToList(root.Right)...)
	return
}
