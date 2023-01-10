package main

// 需要比较左右子树中深度最小的
// 最小深度取决于任一非空子树的长度
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil { // 左右子树为空
		return 1
	}
	if root.Left == nil { // 左子树为空，右子树不为空
		return minDepth(root.Right) + 1
	}
	if root.Right == nil { // 右子树为空，左子树不为空
		return minDepth(root.Left) + 1
	}
	// 左右子树都不为空
	return min(minDepth(root.Left), minDepth(root.Right)) + 1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 迭代
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	depth := 0
	que := []*TreeNode{root}
	for len(que) > 0 {
		depth++
		size := len(que)
		for i := 0; i < size; i++ {
			node := que[0]
			que = que[1:]
			if node.Left == nil && node.Right == nil {
				return depth
			}
			if node.Left != nil {
				que = append(que, node.Left)
			}
			if node.Right != nil {
				que = append(que, node.Right)
			}
		}
	}
	return depth
}
