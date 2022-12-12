package main

// 给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。

/*
	删除倒数第n个节点，实际上，是遍历到距离最后一个节点为n步的节点，该节点的后一个节点即为待删除的节点
    1. fast节点遍历到链表末尾
    2. slow节点和fast节点相距为n，即fast先走n步
*/

// fast先走n步，slow再走，直到fast到达最后一个节点，则slow停止在待删除节点的前一个节点

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyNode := &ListNode{Next: head}
	slow, fast := dummyNode, dummyNode
	i := 0 // i 表示fast走的步数
	for fast != nil && fast.Next != nil {
		if i >= n { // fast先走超过n步了，slow可以走，这时候fast和slow之间距离为n
			slow = slow.Next
		}
		fast = fast.Next
		i++
	}
	if i < n { // fast还没走到n步就到了最后一个节点，n大于链表长度，没有需要删除的节点
		return head
	}
	// 删除slow的后一个节点
	if slow.Next != nil {
		slow.Next = slow.Next.Next
	}
	return dummyNode.Next
}

//
//func removeNthFromEnd(head *ListNode, n int) *ListNode {
//	dummyNode := &ListNode{Next: head}
//	slow, fast := dummyNode, dummyNode
//	// fast走n步, fast和slow之间间隔n个节点
//	for n > 0 && fast != nil {
//		n--
//		fast = fast.Next
//	}
//	// 走完n步fast为nil，则n大于链表长度，或链表为空
//	if fast == nil {
//		return head
//	}
//	// fast走到最后一个节点
//	for fast != nil && fast.Next != nil {
//		fast = fast.Next
//		slow = slow.Next
//	}
//	if slow.Next != nil {
//		slow.Next = slow.Next.Next
//	}
//	return head
//}
