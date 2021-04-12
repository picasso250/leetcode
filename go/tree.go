package main

//   Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	v, _, _ := isValidBSTIter(root)
	return v
}

const INT_MAX = int(^uint(0) >> 1)
const INT_MIN = ^INT_MAX

func isValidBSTIter(root *TreeNode) (isValid bool, max, min int) {
	if root == nil {
		return true, INT_MIN, INT_MAX
	}
	validLeft, leftMax, leftMin := isValidBSTIter(root.Left)
	validRight, rightMax, rightMin := isValidBSTIter(root.Right)
	isValid = validLeft && validRight &&
		root.Val > leftMax && root.Val < rightMin
	max = maxInts(leftMax, rightMax, root.Val)
	min = minInts(leftMin, rightMin, root.Val)
	return
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
