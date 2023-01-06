package main

// 广度优先
func rightSideView(root *TreeNode) []int {
	ans := []int{}
	if root == nil {
		return ans
	}
	que := []*TreeNode{root}
	for len(que) > 0 {
		size := len(que)
		for i := 0; i < size; i++ {
			node := que[0]
			que = que[1:]
			if i == size-1 {
				ans = append(ans, node.Val)
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

// 深度优先 中右左遍历  先遍历右节点，对于每层，只保存最右侧节点
func rightSideView1(root *TreeNode) []int {
	ans := []int{}
	var traversal func(node *TreeNode, depth int)
	traversal = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}
		if depth == len(ans) {
			ans = append(ans, node.Val)
		}
		traversal(node.Right, depth+1) // 先遍历右节点
		traversal(node.Left, depth+1)
	}
	traversal(root, 0)
	return ans
}
