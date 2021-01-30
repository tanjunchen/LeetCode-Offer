package main

// 从尾到头打印链表
// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	if head == nil {
		return nil
	}
	var values []int
	for head != nil {
		values = append(values, head.Val)
		head = head.Next
	}
	start, end := 0, len(values)-1
	for start < end {
		values[start], values[end] = values[end], values[start]
		start++
		end--
	}
	return values
}

func reversePrint2(head *ListNode) []int {
	if head == nil {
		return nil
	}
	var newHead *ListNode
	var res []int
	for head != nil {
		node := head.Next
		head.Next = newHead
		newHead = head
		head = node
	}
	for newHead != nil {
		res = append(res, newHead.Val)
		newHead = newHead.Next
	}
	return res
}
