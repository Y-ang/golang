package main

// 后序遍历：要通过递归函数的返回值来累加求取左叶子数值之和
func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	val := 0
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		val = root.Left.Val
	}
	return sumOfLeftLeaves(root.Left) + sumOfLeftLeaves(root.Right) + val
}

// bfs
func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ans := 0
	que := []*TreeNode{root}
	for len(que) > 0 {
		node := que[0]
		que = que[1:]
		if node.Left != nil && node.Left.Left == nil && node.Left.Right == nil {
			ans += node.Left.Val
		}
		if node.Left != nil {
			que = append(que, node.Left)
		}
		if node.Right != nil {
			que = append(que, node.Right)
		}
	}
	return ans
}
