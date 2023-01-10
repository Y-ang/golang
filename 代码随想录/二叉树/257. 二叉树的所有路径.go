package main

import "strconv"

type TreeNode struct {
	Val   int
	Right *TreeNode
	Left  *TreeNode
}

// dfs
func binaryTreePaths(root *TreeNode) []string {
	ans := []string{}
	var traversal func(node *TreeNode, str string)
	traversal = func(node *TreeNode, str string) {
		if node == nil {
			return
		}
		str += strconv.Itoa(node.Val)
		if node.Left == nil && node.Right == nil {
			ans = append(ans, str)
		}
		traversal(node.Left, str+"->")
		traversal(node.Right, str+"->")
	}
	traversal(root, "")
	return ans
}

// bfs
func binaryTreePaths1(root *TreeNode) []string {
	ans := []string{}
	if root == nil {
		return ans
	}
	stack := []*TreeNode{root}
	paths := []string{strconv.Itoa(root.Val)}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		path := paths[len(paths)-1]
		paths = paths[:len(paths)-1]

		if node.Left == nil && node.Right == nil {
			ans = append(ans, path)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
			paths = append(paths, path+"->"+strconv.Itoa(node.Right.Val))
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
			paths = append(paths, path+"->"+strconv.Itoa(node.Left.Val))
		}

	}
	return ans
}

func main() {
	node1 := &TreeNode{Val: 1}
	node2 := &TreeNode{Val: 2}
	node3 := &TreeNode{Val: 3}
	node5 := &TreeNode{Val: 5}
	node1.Left = node2
	node1.Right = node3
	node2.Left = node5
	binaryTreePaths1(node1)
}
