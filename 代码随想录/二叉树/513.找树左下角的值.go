package main

// bfs
func findBottomLeftValue(root *TreeNode) int {

	maxDepth := -1
	ans := -1
	var traversal func(node *TreeNode, depth int)
	traversal = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}
		if depth > maxDepth { // 记录第一个深度最大的node
			ans = node.Val
			maxDepth = depth
		}
		traversal(node.Left, depth+1)
		traversal(node.Right, depth+1)
	}
	traversal(root, 0)
	return ans
}

// bfs
func findBottomLeftValue(root *TreeNode) int {
	que := []*TreeNode{root}
	ans := 0
	for len(que) > 0 {
		size := len(que)
		for i := 0; i < size; i++ {
			node := que[0]
			que = que[1:]
			if i == 0 { // 记录每层的第一个节点
				ans = node.Val
			}
			if node.Left != nil {
				que = append(que, node.Left)
			}
			if node.Right != nil {
				que = append(que, node.Right)
			}
		}
	}
	return ans
}
