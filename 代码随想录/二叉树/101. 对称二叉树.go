package main

// 递归
// 判断左右子树是否对称，左子树的左节点和右子树的右节点 左子树的右节点和右子树的左节点
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	// 成对比较
	var traversal func(left *TreeNode, right *TreeNode) bool
	traversal = func(left *TreeNode, right *TreeNode) bool {
		// 返回条件
		if left == nil && right == nil { // 左右都为空
			return true
		} else if left == nil || right == nil { // 左右其中一个为空
			return false
		} else if left.Val != right.Val { // 左右节点值不相同
			return false
		}
		// 左右孩子值相同，处理单层
		return traversal(left.Left, right.Right) && traversal(left.Right, right.Left)
	}
	return traversal(root.Left, root.Right)
}

// 迭代 使用队列的方法：把左右两个子树要比较的元素顺序放进一个容器，然后成对成对的取出来进行比较
func isSymmetric1(root *TreeNode) bool {
	if root == nil {
		return true
	}
	que := []*TreeNode{root.Left, root.Right}
	for len(que) > 0 {
		left := que[0] // 成对取出镜像节点
		right := que[1]
		que = que[2:]
		if left == nil && right == nil {
			continue
		} else if left == nil || right == nil || left.Val != right.Val {
			return false
		}
		que = append(que, left.Left) // 比较左子树左节点和右子树右节点
		que = append(que, right.Right)

		que = append(que, left.Right) // 比较左子树右节点和右子树左节点
		que = append(que, right.Left)
	}
	return true
}
