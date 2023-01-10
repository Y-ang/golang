package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(nums []int, index int) *TreeNode {
	if index >= len(nums) || nums[index] == -1 {
		return nil
	}
	node := &TreeNode{Val: nums[index]}
	node.Left = buildTree(nums, 2*index+1)
	node.Right = buildTree(nums, 2*index+2)
	return node
}

func main() {
	nums := []int{3, 9, 20, -1, -1, 15, 7}
	root := buildTree(nums, 0)
	print("%x", root)
}
