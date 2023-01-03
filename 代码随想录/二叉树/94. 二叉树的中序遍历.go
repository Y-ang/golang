package main

// 左中右
//func inorderTraversal(root *TreeNode) []int {
//	var ans []int
//	var traverse func(node *TreeNode)
//	traverse = func(node *TreeNode) {
//		if node == nil {
//			return
//		}
//		traverse(node.Left)
//		ans = append(ans, node.Val)
//		traverse(node.Right)
//	}
//	traverse(root)
//	return ans
//}

func inorderTraversal(root *TreeNode) []int {
	var ans []int
	var stack []*TreeNode
	for root != nil || len(stack) > 0 {
		// 遍历到最左侧节点
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		index := len(stack) - 1 // 栈顶元素
		root = stack[index]
		stack = stack[:index]
		ans = append(ans, root.Val) // 左
		root = root.Right           // 右
	}
	return ans
}
