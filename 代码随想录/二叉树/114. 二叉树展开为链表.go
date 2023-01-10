package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
1、先利用 flatten(x.left) 和 flatten(x.right) 将 x 的左右子树拉平。
2. 将左子树作为右子树
3. 将原先的右子树接到右子树的末端
*/
func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	// 1. 分别flatten左右子树
	flatten(root.Left)
	flatten(root.Right)

	right := root.Right
	// 2. 左子树置空，右子树接上原左子树
	root.Right, root.Left = root.Left, nil
	// 3. 将原右子树接到树的末端
	for root.Right != nil {
		root = root.Right
	}
	root.Right = right

	return
}

/*
1. 将左子树作为右子树
2. 将原右子树放到原左子树的末端
3. 继续遍历右子树的跟节点，知道右子树为nil
*/
func flatten1(root *TreeNode) {
	for root != nil {
		if root.Left != nil {
			right := root.Right
			root.Right, root.Left = root.Left, nil // 左子树移动到右子树, 左子树置空
			p := root
			for p.Right != nil {
				p = p.Right
			}
			p.Right = right // 原右子树接到树的最右端
		}
		root = root.Right // 继续遍历右子树
	}
}

func main() {
	node1 := &TreeNode{Val: 1}
	node2 := &TreeNode{Val: 2}
	node3 := &TreeNode{Val: 3}
	node4 := &TreeNode{Val: 4}
	node5 := &TreeNode{Val: 5}
	node6 := &TreeNode{Val: 6}
	node1.Left = node2
	node1.Right = node5
	node2.Left = node3
	node2.Right = node4
	node5.Right = node6
	flatten1(node1)
}
