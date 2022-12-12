package main

/*
slow 和 fast在环内相遇  ( x + y ) * 2 = x + y + n (y + z )  ===>>  x = (n - 1) (y + z) + z

n为1，意味着fast指针在环形里转了一圈之后，就遇到了 slow指针了。
n大于1，fast指针在环形转n圈之后才遇到 slow指针。
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow { // 如果在环内相遇
			slow = head // 头节点和相遇节点再次相遇，即为环的入口节点
			for fast != slow {
				fast = fast.Next
				slow = slow.Next
			}
			return fast
		}
	}
	return nil
}

func main() {
	tail := &ListNode{-4, nil}
	node1 := &ListNode{0, tail}
	node2 := &ListNode{2, node1}
	head := &ListNode{3, node2}
	detectCycle(head)
}
