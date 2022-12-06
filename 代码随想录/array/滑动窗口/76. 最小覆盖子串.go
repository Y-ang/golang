package main

/*
给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。
注意：
对于 t 中重复字符，我们寻找的子字符串中该字符数量必须不少于 t 中该字符数量。
如果 s 中存在这样的子串，我们保证它是唯一的答案。

输入：s = "ADOBECODEBANC", t = "ABC"
输出："BANC"


滑动窗口：
先统计 t 中每种字符的个数 need := map[string]int{}
窗口的形式遍历 s，定义一个valid表示窗口内覆盖到 t 中字符的个数，当valid >= len(need), 缩小窗口
*/

func minWindow(s string, t string) string {
	var ans string
	minLen := len(s) + 1
	left, valid := 0, 0
	need, win := map[byte]int{}, map[byte]int{}
	// 统计t中字符次数
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}
	for right := 0; right < len(s); right++ {
		rCh := s[right]
		if _, ok := need[rCh]; ok { // win中覆盖need的有效字符数
			win[rCh]++
			if win[rCh] == need[rCh] {
				valid++
			}
		}
		for valid >= len(need) {
			if right-left+1 < minLen {
				ans = s[left : right+1]
				minLen = right - left + 1
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
