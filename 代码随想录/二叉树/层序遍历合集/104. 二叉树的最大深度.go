package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// dfs
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

// bfs
func maxDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ans := 0
	que := []*TreeNode{root}
	for len(que) > 0 {
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
		ans++
	}
	return ans
}
func main() {
	node1 := &TreeNode{Val: 3}
	node2 := &TreeNode{Val: 9}
	node3 := &TreeNode{Val: 20}
	node4 := &TreeNode{Val: 15}
	node5 := &TreeNode{Val: 7}
	node1.Left = node2
	node1.Right = node3
	node3.Left = node4
	node3.Right = node5
	maxDepth1(node1)
}
