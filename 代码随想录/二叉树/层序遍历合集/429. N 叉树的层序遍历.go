package main

type Node struct {
	Val      int
	Children []*Node
}

// 广度优先
func levelOrder(root *Node) [][]int {
	ans := [][]int{}
	if root == nil {
		return ans
	}
	que := []*Node{root}
	for len(que) > 0 {
		size := len(que)
		level := len(ans)
		ans = append(ans, []int{})
		for i := 0; i < size; i++ {
			node := que[0]
			que = que[1:]
			ans[level] = append(ans[level], node.Val)
			for _, v := range node.Children {
				if v != nil {
					que = append(que, v)
				}
			}
		}
	}
	return ans
}

// 深度优先
func levelOrder1(root *Node) [][]int {
	ans := [][]int{}
	var traversal func(node *Node, depth int)
	traversal = func(node *Node, depth int) {
		if node == nil {
			return
		}
		if depth == len(ans) {
			ans = append(ans, []int{})
		}
		ans[depth] = append(ans[depth], node.Val)
		for _, v := range node.Children {
			traversal(v, depth+1)
		}
	}
	traversal(root, 0)
	return ans
}
