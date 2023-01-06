package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 迭代法
func levelOrderBottom(root *TreeNode) [][]int {
	ans := [][]int{}
	if root == nil {
		return ans
	}
	que := []*TreeNode{root} // 保存节点的队列
	for len(que) > 0 {
		level := len(ans)          // 当前层
		ans = append(ans, []int{}) // 添加保存当前层节点值的切片
		size := len(que)           // 当前层节点的个数
		for i := 0; i < size; i++ {
			node := que[0] // 队列头节点
			que = que[1:]
			ans[level] = append(ans[level], node.Val) // 将节点值加入当前层切片
			if node.Left != nil {
				que = append(que, node.Left)
			}
			if node.Right != nil {
				que = append(que, node.Right)
			}
		}
	}
	for i, j := 0, len(ans)-1; i < j; i, j = i+1, j-1 {
		ans[i], ans[j] = ans[j], ans[i]
	}
	return ans
}

// 递归法
func levelOrderBottom1(root *TreeNode) [][]int {
	ans := [][]int{}
	depth := 0
	var traversal func(node *TreeNode, depth int)
	traversal = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}
		if depth == len(ans) {
			ans = append(ans, []int{})
		}
		ans[depth] = append(ans[depth], node.Val)
		traversal(node.Left, depth+1)
		traversal(node.Right, depth+1)
	}
	traversal(root, depth)
	for i, j := 0, len(ans)-1; i < j; i, j = i+1, j-1 {
		ans[i], ans[j] = ans[j], ans[i]
	}
	return ans
}
