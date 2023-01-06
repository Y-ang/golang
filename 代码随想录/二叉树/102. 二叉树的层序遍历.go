package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 迭代写法
func levelOrder(root *TreeNode) [][]int {
	ans := [][]int{}
	if root == nil {
		return ans
	}
	que := []*TreeNode{root}

	for len(que) > 0 {
		level := len(ans)
		ans = append(ans, []int{}) // 添加一维空切片保存当前层的节点
		size := len(que)
		for i := 0; i < size; i++ {
			node := que[0]                            // 队列头节点
			que = que[1:]                             // 弹出队列头节点
			ans[level] = append(ans[level], node.Val) // 将节点值append到当前level层
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

// 递归写法   递归函数中使用depth控制当前节点所在的层数
func levelOrder1(root *TreeNode) [][]int {
	ans := [][]int{}
	depth := 0
	var traversal func(node *TreeNode, depth int)
	traversal = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}
		// 对当前节点的处理
		if depth == len(ans) { // 当前节点的深度大于ans的index范围
			ans = append(ans, []int{})
		}
		ans[depth] = append(ans[depth], node.Val)
		traversal(node.Left, depth+1)
		traversal(node.Right, depth+1)
	}
	traversal(root, depth)
	return ans
}
