package main

// dfs
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil && root.Val == targetSum {
		return true
	}
	return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)
}

// bfs
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	stackNode := []*TreeNode{root}
	stackValue := []int{root.Val}
	for len(stackNode) > 0 {
		node := stackNode[len(stackNode)-1]
		stackNode = stackNode[:len(stackNode)-1]
		sum := stackValue[len(stackValue)-1]
		stackValue = stackValue[:len(stackValue)-1]
		if node.Left == nil && node.Right == nil && sum == targetSum {
			return true
		}
		if node.Left != nil {
			stackNode = append(stackNode, node.Left)
			stackValue = append(stackValue, sum+node.Left.Val)
		}
		if node.Right != nil {
			stackNode = append(stackNode, node.Right)
			stackValue = append(stackValue, sum+node.Right.Val)
		}
	}
	return false
}
