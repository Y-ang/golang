package main

/*
给定两个字符串 s 和 p，找到 s 中所有 p 的 异位词 的子串，返回这些子串的起始索引。不考虑答案输出的顺序。
异位词 指由相同字母重排列形成的字符串（包括相同的字符串）。

示例 1:
输入: s = "cbaebabacd", p = "abc"
输出: [0,6]
解释:
起始索引等于 0 的子串是 "cba", 它是 "abc" 的异位词。
起始索引等于 6 的子串是 "bac", 它是 "abc" 的异位词。

*/

func findAnagrams(s string, p string) []int {
	var ans []int
	left, valid := 0, 0
	need, win := map[byte]int{}, map[byte]int{}
	for i := 0; i < len(p); i++ {
		need[p[i]]++
	}
	for right := 0; right < len(s); right++ {
		rCh := s[right]
		if _, ok := need[rCh]; ok {
			win[rCh]++
			if win[rCh] == need[rCh] {
				valid++
			}
		}
		for right-left+1 >= len(p) {
			if valid == len(need) {
				ans = append(ans, left)
			}
			lCh := s[left]
			left++
			if _, ok := need[lCh]; ok {
				if win[lCh] == need[lCh] {
					valid--
				}
				win[lCh]--
			}
		}
	}
	return ans
}
