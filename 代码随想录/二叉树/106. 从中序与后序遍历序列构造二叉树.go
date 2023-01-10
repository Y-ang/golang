package main

// 根据后序遍历最后一个节点确定中序遍历的根节点
// 中序遍历根节点 的 右子树的根节点：后续遍历的最后一个节点
// 遍历完中序遍历的右子树后，后序遍历的最后一个节点为中序遍历根节点 的 左子树的根节点
// 后序遍历从后向前：中 右 左
// 配合中序遍历序列，判断结束位置
func buildTree(inorder []int, postorder []int) *TreeNode {
	valueToIndex := map[int]int{} // 构建inoder中value-index的map，便于根据postorder值索引inder下标
	for index, value := range inorder {
		valueToIndex[value] = index
	}
	postIndex := len(postorder) - 1 // 按照postoder从后向前的顺序构建树

	var traversal func(left, right int) *TreeNode
	traversal = func(left, right int) *TreeNode {
		if left > right { // inorder中需要构建树的子数组
			return nil
		}
		val := postorder[postIndex] // 新建后序遍历的最后一个值为新节点
		node := &TreeNode{Val: val}
		inIndex := valueToIndex[val] // 找到该节点在中序遍历中的位置
		postIndex--                  // 每新建一个节点，postorder向前移动一个位置

		node.Right = traversal(inIndex+1, right)
		node.Left = traversal(left, inIndex-1)

		return node
	}
	return traversal(0, len(inorder)-1)
}
