package main

type Node struct {
	Val      int
	Children []*Node
}

// 递归
func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	depth := 0
	for _, v := range root.Children {
		depth = max(depth, maxDepth(v))
	}
	return depth + 1
}

// 迭代
func maxDepth1(root *Node) int {
	if root == nil {
		return 0
	}
	que := []*Node{root}
	depth := 0
	for len(que) > 0 {
		depth++
		size := len(que)
		for i := 0; i < size; i++ {
			node := que[0]
			que = que[1:]
			for _, v := range node.Children {
				if v != nil {
					que = append(que, v)
				}
			}
		}
	}
	return depth
}
