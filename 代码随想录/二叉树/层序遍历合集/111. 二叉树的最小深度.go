package main

// 只有节点的左右孩子都为空，该节点才又可能被记录为最低点
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left != nil && root.Right != nil {
		return min(minDepth(root.Left), minDepth(root.Right)) + 1
	} else if root.Left != nil {
		return minDepth(root.Left) + 1
	} else if root.Right != nil {
		return minDepth(root.Right) + 1
	}
	return 1 // 左右节点都为空
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minDepth1(root *TreeNode) int {
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
