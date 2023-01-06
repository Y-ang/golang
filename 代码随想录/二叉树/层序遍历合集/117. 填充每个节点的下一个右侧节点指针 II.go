package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	que := []*Node{root}
	for len(que) > 0 {
		size := len(que)
		for i := 0; i < size; i++ {
			node := que[0]
			que = que[1:]
			if i < size-1 {
				node.Next = que[0]
			} else {
				node.Next = nil
			}
			if node.Left != nil {
				que = append(que, node.Left)
			}
			if node.Right != nil {
				que = append(que, node.Right)
			}
		}
	}
	return root
}

func main() {
	node1 := &Node{Val: 1}
	node2 := &Node{Val: 2}
	node3 := &Node{Val: 3}
	node4 := &Node{Val: 4}
	node5 := &Node{Val: 5}
	node6 := &Node{Val: 7}
	node1.Left = node2
	node1.Right = node3
	node2.Left = node4
	node2.Right = node5
	node3.Right = node6
	connect(node1)
}

//func connect(root *Node) *Node {
//	var traversal func(node *Node)
//	traversal = func(node *Node) {
//		if node == nil {
//			return
//		}
//		if node.Left != nil {
//			node.Left.Next = node.Right
//		}
//		if node.Right != nil && node.Next != nil {
//			if node.Next.Left != nil {
//				node.Right.Next = node.Next.Left
//			} else {
//				node.Right.Next = node.Next.Right
//			}
//		}
//		traversal(node.Left)
//		traversal(node.Right)
//	}
//	traversal(root)
//	return root
//}
