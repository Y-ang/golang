package main

type TreeNode struct {
	Val   int
	Right *TreeNode
	Left  *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var maxDepth func(node *TreeNode) int
	maxDepth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		return max(maxDepth(node.Left), maxDepth(node.Right)) + 1
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	if abs(left-right) > 1 {
		return false
	}
	return isBalanced(root.Left) && isBalanced(root.Right)
}
func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// 改进：如果左右子树出现任意一个不平衡的树，则返回
func isBalanced1(root *TreeNode) bool {
	return depth(root) != -1
}

// 如果平衡返回当前节点的最大高度，若不平衡返回-1
func depth(node *TreeNode) int {
	if node == nil {
		return 0
	}
	left := depth(node.Left)
	if left == -1 { // 左子树不平衡直接返回-1
		return -1
	}
	right := depth(node.Right)
	if right == -1 || abs(left-right) > 1 { // 右子树不平衡，或左右子树相差大于1即当前子树不平衡
		return -1
	}
	// 当前子树平衡，返回当前子树高度
	return max(left, right) + 1
}

func main() {
	isBalanced(&TreeNode{Val: 1})
}
