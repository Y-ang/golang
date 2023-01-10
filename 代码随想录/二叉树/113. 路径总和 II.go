package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	ans := [][]int{}
	var traversal func(node *TreeNode, targetSum int, pathSeq []int)
	traversal = func(node *TreeNode, targetSum int, pathSeq []int) {
		if node == nil {
			return
		}
		pathSeq = append(pathSeq, node.Val)
		if node.Left == nil && node.Right == nil && targetSum == node.Val {
			pathCopy := make([]int, len(pathSeq)) // 保证保存到ans中slice的地址空间不同，这样使结果不会被后续操作覆盖数据
			copy(pathCopy, pathSeq)
			ans = append(ans, pathCopy)
		}
		traversal(node.Left, targetSum-node.Val, pathSeq)
		traversal(node.Right, targetSum-node.Val, pathSeq)
	}
	traversal(root, targetSum, nil)
	return ans
}

func main() {
	node1 := &TreeNode{Val: 5}
	node2 := &TreeNode{Val: 4}
	node3 := &TreeNode{Val: 8}
	node4 := &TreeNode{Val: 11}
	node5 := &TreeNode{Val: 13}
	node6 := &TreeNode{Val: 4}
	node7 := &TreeNode{Val: 7}
	node8 := &TreeNode{Val: 2}
	node9 := &TreeNode{Val: 5}
	node10 := &TreeNode{Val: 1}
	node1.Left = node2
	node1.Right = node3
	node2.Left = node4
	node3.Left = node5
	node3.Right = node6
	node4.Left = node7
	node4.Right = node8
	node6.Left = node9
	node6.Right = node10
	pathSum(node1, 22)
}
