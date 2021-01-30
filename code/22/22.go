package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

// 链表中倒数第 k 个节点
func getKthFromEnd(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	slow, fast := head, head
	for k > 0 {
		fast = fast.Next
		k--
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}

func getKthFromEnd2(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	cur := head
	count := 0
	for head != nil {
		head = head.Next
		count++
	}
	for i := 0; i < count-k; i++ {
		cur = cur.Next
	}
	return cur
}
