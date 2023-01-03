package main

// 左右中
func postorderTraversal(root *TreeNode) []int {
	var ans []int
	var stack []*TreeNode
	if root != nil {
		stack = append(stack, root)
	}
	for len(stack) > 0 {
		index := len(stack) - 1
		node := stack[index]
		stack = stack[:index]

		ans = append(ans, node.Val)
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}
	reverse(ans)
	return ans
}

func reverse(ans []int) {
	for left, right := 0, len(ans)-1; left < right; {
		ans[left], ans[right] = ans[right], ans[left]
		left, right = left+1, right-1
	}
}

//func postorderTraversal(root *TreeNode) []int {
//	var ans []int
//	var traverse func(node *TreeNode)
//	traverse = func(node *TreeNode) {
//		if node == nil {
//			return
//		}
//		traverse(node.Left)
//		traverse(node.Right)
//		ans = append(ans, node.Val)
//	}
//	traverse(root)
//	return ans
//}
