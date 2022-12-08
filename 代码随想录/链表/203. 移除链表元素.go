package main

/*
删除节点：
判断当前节点的下一个节点是否为待删除节点，如果是让节点next指针直接指向下下一个节点
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	dummyHead := &ListNode{}
	dummyHead.Next = head
	cur := dummyHead
	for cur != nil && cur.Next != nil {
		if cur.Next.Val == val { // cur.Next指向的节点为待删除节点
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next // cur.Next指向的节点非待删除节点
		}

	}
	return dummyHead.Next
}
