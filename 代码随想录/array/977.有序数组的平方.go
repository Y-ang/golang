package main

/*
给你一个按 非递减顺序 排序的整数数组 nums，返回 每个数字的平方 组成的新数组，要求也按 非递减顺序 排序。

示例 1：
输入：nums = [-4,-1,0,3,10]
输出：[0,1,9,16,100]
解释：平方后，数组变为 [16,1,0,9,100]
排序后，数组变为 [0,1,9,16,100]
示例 2：
输入：nums = [-7,-3,2,3,11]
输出：[4,9,9,49,121]

两边大中间小
*/

func sortedSquares(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	left, right := 0, n-1
	for pos := n - 1; pos >= 0; pos-- {
		if v, w := nums[left]*nums[left], nums[right]*nums[right]; v > w {
			ans[pos] = v
			left++
		} else {
			ans[pos] = w
			right--
		}
	}
	return ans
}
