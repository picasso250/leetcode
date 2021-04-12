package main

import (
	"fmt"
	"strconv"
	"strings"
)

//  * Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	p := head
	var prevDel *ListNode = nil
	var toDel *ListNode = nil
	var prev *ListNode
	i := 0
	for p != nil {
		if i >= n-1 {
			prevDel = prev
		}
		if i >= n {
			toDel = p
		}
		i++
		prev = p
		p = p.Next
	}
	// del
	head = head.delNode(toDel, prevDel)
	return head
}
func (head *ListNode) delNode(toDel *ListNode, prevDel *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if toDel == head {
		newHead := head.Next
		head.Next = nil
		return newHead
	}
	if prevDel != nil {
		if prevDel.Next != toDel {
			panic("prevDel's next is not toDel")
		}
		prevDel.Next = toDel.Next
		toDel.Next = nil
	}
	return head
}
func toList(lst []int) *ListNode {
	if len(lst) == 0 {
		return nil
	}
	l := make([]ListNode, len(lst))
	for i, v := range lst {
		var next *ListNode
		if i == len(lst)-1 {
			next = nil
		} else {
			next = &l[i+1]
		}
		l[i] = ListNode{v, next}
	}
	return &l[0]
}
func (head *ListNode) toSlice() []int {
	a := make([]int, 0)
	for p := head; p != nil; p = p.Next {
		a = append(a, p.Val)
	}
	return a
}
func (head *ListNode) toStrings() []string {
	a := make([]string, 0)
	for p := head; p != nil; p = p.Next {
		a = append(a, strconv.Itoa(p.Val))
	}
	return a
}
func (head *ListNode) ToString() string {
	return fmt.Sprintf("[%s]", strings.Join(head.toStrings(), ","))
}
func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{0, head}
	for p := dummy; p != nil; {
		if p.Next != nil && p.Next.Next != nil && p.Next.Val == p.Next.Next.Val {
			p = deleteDuplicatesN(p, p.Next.Val)
		} else {
			p = p.Next
		}
	}
	return dummy.Next
}
func deleteDuplicatesN(dummy *ListNode, val int) *ListNode {
	for p := dummy; p.Next != nil; {
		if p.Next.Val == val {
			p.Next = p.Next.Next
		} else {
			break
		}
	}
	return dummy
}
