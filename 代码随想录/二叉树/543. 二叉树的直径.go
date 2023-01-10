package main

// 每一条二叉树的「直径」长度，就是一个节点的左右子树的最大深度之和
// 前序位置无法获取子树信息，所以只能让每个节点调用 maxDepth 函数去算子树的深度
// 把计算「直径」的逻辑放在后序位置，准确说应该是放在 maxDepth 的后序位置，因为 maxDepth 的后序位置是知道左右子树的最大深度的。
func diameterOfBinaryTree(root *TreeNode) int {
	maxDia := 0
	var maxDepth func(root *TreeNode) int
	maxDepth = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		leftDepth := maxDepth(root.Left)
		rightDepth := maxDepth(root.Right)
		// 计算以当前节点为根的子树的最大直径
		// 后序位置，顺便计算最大直径
		maxDia = max(maxDia, leftDepth+rightDepth)
		return max(leftDepth, rightDepth) + 1
	}
	maxDepth(root)
	return maxDia
}
