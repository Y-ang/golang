package main

type MyLinkedList struct {
	head *Node
}

type Node struct {
	Val  int
	Prev *Node
	Next *Node
}

func Constructor() MyLinkedList {
	return MyLinkedList{}
}

// 遍历index步，若节点不为空返回当前节点值；若为空返回-1
func (this *MyLinkedList) Get(index int) int {
	cur := this.head
	for i := 0; i < index && cur != nil; i++ {
		cur = cur.Next
	}
	if cur != nil {
		return cur.Val
	}
	return -1
}

// 新node的下一个节点指向head，若head不为空head的前一个节点指向node
func (this *MyLinkedList) AddAtHead(val int) {
	node := &Node{
		Val:  val,
		Next: this.head,
	}
	if this.head != nil {
		this.head.Prev = node
	}
	this.head = node
}

// 若head为空，等价于AddAtHead；若head不为空，遍历到最后一个节点插入
func (this *MyLinkedList) AddAtTail(val int) {
	if this.head == nil {
		this.AddAtHead(val)
	} else {
		node := &Node{Val: val}
		cur := this.head
		for cur != nil && cur.Next != nil {
			cur = cur.Next
		}
		cur.Next = node
		node.Prev = cur
	}
}

// 若index=0在头插入，否则遍历index - 1步然后执行插入操作
func (this *MyLinkedList) AddAtIndex(index int, val int) { // 找到index - 1位置执行插入操作
	if index == 0 {
		this.AddAtHead(val)
	} else {
		cur := this.head
		for i := 0; i < index-1 && cur != nil; i++ {
			cur = cur.Next
		}
		if cur != nil { // 防止index大于链表长度的非法情况的出现
			node := &Node{Val: val}
			node.Prev = cur
			node.Next = cur.Next

			if cur.Next != nil {
				cur.Next.Prev = node
			}

			cur.Next = node
		}
	}
}

// 若index = 0删除头节点，否则遍历index - 1步然后执行插入操作
func (this *MyLinkedList) DeleteAtIndex(index int) { // 找到index - 1位置执行删除操作
	if index == 0 {
		this.head = this.head.Next
		if this.head != nil {
			this.head.Prev = nil
		}
	} else {
		cur := this.head
		for i := 0; i < index-1 && cur != nil; i++ {
			cur = cur.Next
		}
		if cur != nil && cur.Next != nil {
			cur.Next = cur.Next.Next
			if cur.Next != nil {
				cur.Next.Prev = cur
			}
		}
	}
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
func main() {
	obj := Constructor()
	obj.AddAtHead(1)
	obj.AddAtTail(3)
	obj.AddAtIndex(1, 2)
	obj.DeleteAtIndex(1)
}
