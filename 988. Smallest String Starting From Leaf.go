package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func smallestFromLeaf(root *TreeNode) string {
	return to(s(root))
}
func s(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}
	if root.Left != nil && root.Right != nil {
		a := append(s(root.Left),root.Val)
		b := append(s(root.Right),root.Val)
		if cmp(a, b) {
			return b
		}
		return a
	}
	if root.Left == nil {
		return append(s(root.Right), root.Val)
	}
	return append(s(root.Left), root.Val)
}
func to(a []int) string {
	b := make([]byte, len(a))
	for i, c := range a {
		b[i] = byte(c) + 'a'
	}
	return string(b)
}
func cmp(a []int, b []int) bool {
	fmt.Printf("cmp %v %v\n",a,b)
	la := len(a)
	lb := len(b)
	min := la
	if la > lb {
		min = lb
	}
	for i := 0; i < min; i++ {
		if a[i] > b[i] {
			return true
		}
		if a[i] < b[i] {
			return false
		}
	}
	return la > lb
}
func main() {
	t := &TreeNode{
		0,
		&TreeNode{
			1,
			&TreeNode{
				3,
				nil,
				nil,
			},
			&TreeNode{
				4,
				nil,
				nil,
			},
		},
		&TreeNode{
			2,
			&TreeNode{
				3,
				nil,
				nil,
			},
			&TreeNode{
				4,
				nil,
				nil,
			},
		},
	}
	t = &TreeNode{
		4,
		&TreeNode{
			0,
			&TreeNode{
				1,
				nil,
				nil,
			},
			nil,
		},
		&TreeNode{
			1,
			nil,
			nil,
		},
	}
	t = &TreeNode{
		25,
		&TreeNode{
			1,
			&TreeNode{
				0,
				&TreeNode{
					1,
					&TreeNode{
						0,
						nil,
						nil,
					},
					nil,
				},
				nil,
			},
			&TreeNode{
				0,
				nil,
				nil,
			},
		},
		nil,
	}
	fmt.Printf("%s\n", smallestFromLeaf(t))
}
