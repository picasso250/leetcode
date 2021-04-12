package main

import "fmt"

//    Definition for a binary tree node.
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

// func treeToList(root *TreeNode) (ret []int) {
// 	if root == nil {
// 		return
// 	}
// 	ret = []int{root.Val}
// 	ret = append(ret, treeToList(root.Left)...)
// 	ret = append(ret, treeToList(root.Right)...)
// 	return
// }
func recoverTree(root *TreeNode) {
	if root == nil {
		return
	}
	lst := midOrderPointer(root)
	i, j := -1, -1
	for k := 0; k < len(lst)-1; k++ {
		if lst[k].Val > lst[k+1].Val {
			if i == -1 {
				i = k
			} else {
				j = k
			}
		}
	}

	// mapParent := make(map[*TreeNode]*TreeNode)
	// mapDirection := make(map[*TreeNode]bool)
	// treeInfo(root, mapParent, mapDirection)
	a, b := lst[i], lst[j+1]
	if j == -1 {
		a, b = lst[i], lst[i+1]
	}
	a.Val, b.Val = b.Val, a.Val
}

func midOrderPointer(root *TreeNode) (ret []*TreeNode) {
	if root == nil {
		return
	}
	ret = append(ret, midOrderPointer(root.Left)...)
	ret = append(ret, root)
	ret = append(ret, midOrderPointer(root.Right)...)
	return
}

const TreeLeft = true
const TreeRight = false

func treeInfo(root *TreeNode,
	mapParent map[*TreeNode]*TreeNode,
	mapDirection map[*TreeNode]bool) {
	if root == nil {
		return
	}
	mapParent[root.Left] = root
	mapParent[root.Right] = root
	mapDirection[root.Left] = TreeLeft
	mapDirection[root.Right] = TreeRight
	treeInfo(root.Left, mapParent, mapDirection)
	treeInfo(root.Right, mapParent, mapDirection)
}
func buildTree(a []interface{}) *TreeNode {
	n := len(a)
	if n == 0 {
		return nil
	}
	ret := make([]*TreeNode, n)
	for i, v := range a {
		if v == nil {
			ret[i] = nil
		} else {
			ret[i] = &TreeNode{v.(int), nil, nil}
		}
	}
	for i, v := range ret {
		if v == nil {

		} else {
			left := i*2 + 1
			right := i*2 + 2
			if left < n {
				v.Left = ret[left]
			}
			if right < n {
				v.Right = ret[right]
			}
		}
	}
	return ret[0]
}
func treeToList(tree *TreeNode) []interface{} {
	if tree == nil {
		return []interface{}{}
	}
	i, j := 0, 1
	ret := []*TreeNode{tree}
	for {
		notAllNil := false
		for k := i; k < j; k++ {
			if ret[k] != nil {
				// ret[2*k+1] = ret[k].Left
				// ret[2*k+2] = ret[k].Right
				ret = append(ret, ret[k].Left, ret[k].Right)
				if ret[k].Left != nil || ret[k].Right != nil {
					notAllNil = true
				}
			}
		}
		i = 2*i + 1
		j = 2*(i+1) + 1
		if !notAllNil {
			break
		}
	}
	j = len(ret) - 1
	for j >= 0 && ret[j] == nil {
		j--
	}
	j++
	r := make([]interface{}, j)
	for i = 0; i < j; i++ {
		if ret[i] == nil {
			r[i] = nil
		} else {
			r[i] = ret[i].Val
		}
	}
	return r
}
func treeToString(tree *TreeNode) string {
	if tree == nil {
		return "nil"
	}
	return fmt.Sprintf("(%s %d %s)",
		treeToString(tree.Left), tree.Val, treeToString(tree.Right))
}
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return absInt(treeHeight(root.Left)-treeHeight(root.Right)) <= 1 &&
		isBalanced(root.Left) && isBalanced(root.Right)
}
func treeHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return maxInts(treeHeight(root.Left), treeHeight(root.Right)) + 1
}
func isBalancedIter(root *TreeNode) (isBal bool, height int) {
	if root == nil {
		return true, 0
	}
	leftIsBal, leftHeight := isBalancedIter(root.Left)
	rightIsBal, rightHeight := isBalancedIter(root.Right)
	isBal = absInt(leftHeight-rightHeight) <= 1 && leftIsBal && rightIsBal
	height = maxInts(leftHeight, rightHeight) + 1
	return
}
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	if root.Left == nil {
		return minDepth(root.Right) + 1
	}
	if root.Right == nil {
		return minDepth(root.Left) + 1
	}
	return minInts(minDepth(root.Left), minDepth(root.Right)) + 1
}
