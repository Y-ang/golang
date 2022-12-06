package main

/*
给你一个正整数 n ，生成一个包含 1 到 n2 所有元素，且元素按顺时针顺序螺旋排列的 n x n 正方形矩阵 matrix 。
输入：n = 3
输出：[[1,2,3],[8,9,4],[7,6,5]]

模拟顺时针画矩阵的过程:
填充上行从左到右
填充右列从上到下
填充下行从右到左
填充左列从下到上
由外向内一圈一圈这么画下去。
*/

func generateMatrix(n int) [][]int {
	left, right, top, bottom := 0, n-1, 0, n-1
	number := 1

	ans := make([][]int, n)
	for i, _ := range ans {
		ans[i] = make([]int, n)
	}

	for number <= n*n {
		for j := left; j <= right; j++ {
			ans[top][j] = number
			number++
		}
		top++
		for i := top; i <= bottom; i++ {
			ans[i][right] = number
			number++
		}
		right--
		for j := right; j >= left; j-- {
			ans[bottom][j] = number
			number++
		}
		bottom--
		for i := bottom; i >= top; i-- {
			ans[i][left] = number
			number++
		}
		left++
	}
	return ans
}
