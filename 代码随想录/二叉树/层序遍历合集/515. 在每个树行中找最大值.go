package main

import "math"

// 广度
func largestValues(root *TreeNode) []int {
	ans := []int{}
	if root == nil {
		return ans
	}
	que := []*TreeNode{root}
	for len(que) > 0 {
		size := len(que)
		max := math.MinInt32
		for i := 0; i < size; i++ {
			node := que[0]
			que = que[1:]
			if node.Val > max {
				max = node.Val
			}
			if node.Left != nil {
				que = append(que, node.Left)
			}
			if node.Right != nil {
				que = append(que, node.Right)
			}
		}
		ans = append(ans, max)
	}
	return ans
}

// 深度
func largestValues1(root *TreeNode) []int {
	ans := []int{}
	var traversal func(node *TreeNode, depth int)
	traversal = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}
		if depth < len(ans) && ans[depth] < node.Val {
			ans[depth] = node.Val
		} else if depth == len(ans) { // 新level
			ans = append(ans, node.Val)
		}
		traversal(node.Left, depth+1)
		traversal(node.Right, depth+1)
	}
	traversal(root, 0)
	return ans
}
