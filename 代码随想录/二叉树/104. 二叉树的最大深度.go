package main

/*
二叉树节点的深度：指从根节点到该节点的最长简单路径边的条数或者节点数（取决于深度从0开始还是从1开始）
二叉树节点的高度：指从该节点到叶子节点的最长简单路径边的条数后者节点数（取决于高度从0开始还是从1开始）
而根节点的高度就是二叉树的最大深度，所以本题中我们通过后序求的根节点高度来求的二叉树最大深度。
*/

// 递归 后序遍历
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 迭代 层序遍历
func maxDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	que := []*TreeNode{root}
	depth := 0
	for len(que) > 0 {
		depth++
		size := len(que)
		for i := 0; i < size; i++ {
			node := que[0]
			que = que[1:]
			if node.Left != nil {
				que = append(que, node.Left)
			}
			if node.Right != nil {
				que = append(que, node.Right)
			}
		}
	}
	return depth
}
