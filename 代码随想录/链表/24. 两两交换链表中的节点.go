package main

// 给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。你必须在不修改节点内部的值的情况下完成本题（即，只能进行节点交换）。

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	dummyNode := &ListNode{Next: head}
	pre := dummyNode
	cur := head
	for cur != nil && cur.Next != nil { // 存在成对的节点待反转
		post := cur.Next.Next

		pre.Next = cur.Next // 步骤一：前面的节点指向当前节点的下一个节点
		cur.Next.Next = cur // 步骤二：当前节点的下一个节点指向当前节点
		cur.Next = post     // 步骤三：当前节点指向post

		pre = cur
		cur = cur.Next
	}
	return dummyNode.Next
}
