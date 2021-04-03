package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func makeList(a ...int) *ListNode {
	var nd *ListNode
	for i := len(a) - 1; i >= 0; i-- {
		nd = &ListNode{a[i], nd}
	}
	return nd
}
func printList(head *ListNode) {
	for ; head != nil; head = head.Next {
		fmt.Printf("%d->", head.Val)
	}
	fmt.Println()
}
func removeElements(head *ListNode, val int) *ListNode {
	var prev *ListNode
	var p *ListNode
	for p = head; p != nil; p = p.Next {
		if p.Val == val {
			if prev == nil { // head
				head = p.Next
			} else {
				prev.Next = p.Next
			}
		} else {
			prev = p
		}
	}
	return head
}
func main() {
	lst := makeList(1, 1)
	lst = removeElements(lst, 1)
	printList(lst)
}
