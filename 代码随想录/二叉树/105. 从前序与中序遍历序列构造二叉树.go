package main

func buildTree(preorder []int, inorder []int) *TreeNode {
	valueToIdx := map[int]int{}
	for index, value := range inorder {
		valueToIdx[value] = index
	}
	preIndex := 0
	var traversal func(left, right int) *TreeNode
	traversal = func(left, right int) *TreeNode {
		if left > right {
			return nil
		}
		value := preorder[preIndex]
		root := &TreeNode{Val: value}
		preIndex++

		InIdx := valueToIdx[value]
		root.Left = traversal(left, InIdx-1)
		root.Right = traversal(InIdx+1, right)
		return root
	}
	return traversal(0, len(inorder)-1)
}

func main() {
	inorder := []int{9, 3, 15, 20, 7}
	preorder := []int{9, 15, 7, 20, 3}
	buildTree(preorder, inorder)

}

//func buildTree(preorder []int, inorder []int) *TreeNode {
//	valueToIdx := map[int]int{}
//	for index, value := range inorder {
//		valueToIdx[value] = index
//	}
//	preIndex := 0
//	var traversal func(left, right int) *TreeNode
//	traversal = func(left, right int) *TreeNode {
//		if left > right {
//			return nil
//		}
//		value := preorder[preIndex]
//		root := &TreeNode{Val: value}
//		preIndex++
//
//		InIdx := valueToIdx[value]
//		root.Left = traversal(left, InIdx-1)
//		root.Right = traversal(InIdx+1, right)
//		return root
//	}
//	return traversal(0, len(inorder)-1)
//}
