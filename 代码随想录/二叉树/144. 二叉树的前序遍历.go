package main

type TreeNode struct {
	Val   int
	Right *TreeNode
	Left  *TreeNode
}

// 中左右
//func preorderTraversal(root *TreeNode) []int {
//	var ans []int
//	var traverse func(node *TreeNode)
//	traverse = func(node *TreeNode) {
//		if node == nil {
//			return
//		}
//		ans = append(ans, node.Val)
//		traverse(node.Left)
//		traverse(node.Right)
//	}
//	traverse(root)
//	return ans
//}

func preorderTraversal(root *TreeNode) []int {
	var ans []int
	var stack []*TreeNode
	if root != nil {
		stack = append(stack, root)
	}
	for len(stack) > 0 {
		index := len(stack) - 1 // 栈顶
		node := stack[index]
		stack = stack[:index] // 栈顶元素出栈

		ans = append(ans, node.Val)
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
	return ans
}

func main() {
	root := &TreeNode{Val: 1}
	node1 := &TreeNode{Val: 2}
	node2 := &TreeNode{Val: 3}
	node1.Left = node2
	root.Left = node1

	preorderTraversal(root)
}
