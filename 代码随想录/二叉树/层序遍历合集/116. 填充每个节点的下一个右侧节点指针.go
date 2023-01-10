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

// 深度
func connect(root *Node) *Node {

	if root == nil || (root.Left == nil && root.Right == nil) {
		return root
	}
	// 对当前节点的处理
	root.Left.Next = root.Right // 左孩子的next指向右孩子
	if root.Next != nil {       // 如果当前节点的next节点存在，则右孩子的next指向当前next节点的左孩子
		root.Right.Next = root.Next.Left
	}
	connect(root.Left)
	connect(root.Right)

	return root
}
