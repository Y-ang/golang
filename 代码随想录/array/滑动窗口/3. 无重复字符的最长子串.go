package main

/*
给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。

示例 1:
输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
*/

func lengthOfLongestSubstring(s string) int {
	ans := 0
	left := 0
	win := map[byte]int{}
	for right := 0; right < len(s); right++ {
		win[s[right]]++
		for win[s[right]] > 1 {
			win[s[left]]--
			left++
		}
		ans = max(ans, right-left+1)
	}
	return ans
}

func max(x, y int) int {
	if y > x {
		return y
	}
	return x
}

func main() {

}
