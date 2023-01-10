package main

// 满二叉树的节点个数：2^h - 1 （h为当前节点的高度）
// 完全二叉树节点个数是普通二叉树和完全二叉树的结合版
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// 满二叉树的节点个数计算
	rootLeft, rootRight := root, root
	leftH := 0
	for rootLeft != nil {
		leftH++
		rootLeft = rootLeft.Left
	}
	rightH := 0
	for rootRight != nil {
		rightH++
		rootRight = rootRight.Right
	}
	// 是满二叉树
	if leftH == rightH {
		return 1<<leftH - 1 // 左移n位就是乘以2的n次方
	}
	// 不是满二叉树，按照普通二叉树的节点计算
	return countNodes(root.Left) + countNodes(root.Right) + 1
}
