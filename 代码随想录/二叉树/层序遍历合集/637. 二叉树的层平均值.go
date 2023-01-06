package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 广度优先
func averageOfLevels(root *TreeNode) []float64 {
	ans := []float64{}
	if root == nil {
		return ans
	}
	que := []*TreeNode{root}
	for len(que) > 0 {
		levelSum := 0
		size := len(que)
		for i := 0; i < size; i++ {
			node := que[0]
			que = que[1:]
			levelSum += node.Val
			if node.Left != nil {
				que = append(que, node.Left)
			}
			if node.Right != nil {
				que = append(que, node.Right)
			}
		}
		ans = append(ans, float64(levelSum)/float64(size))
	}
	return ans
}

type data struct {
	Sum, Cnt int
}

func averageOfLevels1(root *TreeNode) []float64 {
	levelData := []data{}

	var traversal func(node *TreeNode, depth int)
	traversal = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}
		if depth < len(levelData) { // 之前遍历过的层
			levelData[depth].Sum += node.Val
			levelData[depth].Cnt++
		} else { // 新的一层
			levelData = append(levelData, data{Sum: node.Val, Cnt: 1})
		}
		traversal(node.Left, depth+1)
		traversal(node.Right, depth+1)
	}
	traversal(root, 0)
	ans := []float64{}
	for _, v := range levelData {
		ans = append(ans, float64(v.Sum)/float64(v.Cnt))
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
	//averageOfLevels(node1)
	averageOfLevels1(node1)
}
