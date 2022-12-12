package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	for head != nil {
		post := head.Next
		head.Next = pre
		pre = head
		head = post
	}
	return pre
}

// 递归
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var p = reverseList(head.Next) // 反转后面的，并保存最后一个节点
	head.Next.Next = head          // 反转head.Next指针
	head.Next = nil
	return p // 返回最后一个节点
}
