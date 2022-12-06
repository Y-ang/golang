package main

/*
统计一个数字在排序数组中出现的次数。

示例 1:
输入: nums = [5,7,7,8,8,10], target = 8
输出: 2
示例 2:
输入: nums = [5,7,7,8,8,10], target = 6
输出: 0

与34题相同仅返回值不同，找到target的左右边界，返回边界的长度
左边届：target 的左边界
右边界：target + 1 的左边界 - 1
判断nums左右边界对应的值是否为target

注：因为左边界的范围为[0, len(nums)], 所以要判断边界范围，以及查找到的下标对应的值是否为target
*/

func search(nums []int, target int) int {
	searchLeftIndex := func(number int) int {
		left, right := 0, len(nums)
		for left < right {
			mid := left + (right-left)>>1
			if nums[mid] >= number {
				right = mid
			} else {
				left = mid + 1
			}
		}
		return left
	}
	left, right := searchLeftIndex(target), searchLeftIndex(target+1)-1
	if left <= right && right < len(nums) && nums[left] == target && nums[right] == target {
		return right - left + 1
	}
	return 0
}
func main() {
	println(search([]int{5, 7, 7, 8, 8, 10}, 8))
}
