package main

/*
给你两个字符串 s1 和 s2 ，写一个函数来判断 s2 是否包含 s1 的排列。如果是，返回 true ；否则，返回 false 。
换句话说，s1 的排列之一是 s2 的 子串 。

示例 1：
输入：s1 = "ab" s2 = "eidbaooo"
输出：true
解释：s2 包含 s1 的排列之一 ("ba").

*/

func checkInclusion(s1 string, s2 string) bool {
	left, valid := 0, 0
	need, win := map[byte]int{}, map[byte]int{}

	for i := 0; i < len(s1); i++ {
		need[s1[i]]++
	}

	for right := 0; right < len(s2); right++ {
		rCh := s2[right]
		if _, ok := need[rCh]; ok {
			win[rCh]++
			if win[rCh] == need[rCh] {
				valid++
			}
		}
		for right-left+1 >= len(s1) {
			if valid == len(need) {
				return true
			}
			lCh := s2[left]
			left++
			if _, ok := need[lCh]; ok {
				if win[lCh] == need[lCh] {
					valid--
				}
				win[lCh]--
			}
		}
	}

	return false
}
