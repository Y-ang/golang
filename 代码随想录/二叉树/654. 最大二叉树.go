package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) <= 0 {
		return nil
	}
	best := 0
	for i, v := range nums {
		if v > nums[best] {
			best = i
		}
	}
	node := &TreeNode{Val: nums[best]}

	node.Left = constructMaximumBinaryTree(nums[:best])

	node.Right = constructMaximumBinaryTree(nums[best+1:])

	return node
}

func main() {
	nums := []int{3, 2, 1, 6, 0, 5}
	root := constructMaximumBinaryTree(nums)
	print(root)
}
