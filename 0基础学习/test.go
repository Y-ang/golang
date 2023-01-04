package main

import (
	"fmt"
	"unicode"
)

func main() {
	testArr := [5]int{0, 1, 2, 3, 4}
	//reverse(&testArr)
	fmt.Println(testArr)
	a := rotate(testArr[:], 2)
	fmt.Println(a)

	b := []string{"tao", "taoshihan", "shi", "shi", "han"}
	emptyString(b)
	d := []byte("abc bcd wer  sdsd  taoshihan     de")
	e := emptyString2(d)
	fmt.Println(string(e))
	f := []byte("abc bcd wer  sdsd  taoshihan     de")
	reverse1(f)
	fmt.Println(string(f))
}

/*
ç»ä¹  4.3ï¼ éåreverseå½æ°ï¼ä½¿ç¨æ°ç»æéä»£æ¿sliceã
*/
func reverse(s *[5]int) {
	i, j := 0, len(*s)-1
	for i < j {
		(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
		i += 1
		j -= 1
	}
}

/*
ç»ä¹  4.4ï¼ ç¼åä¸ä¸ªrotateå½æ°ï¼éè¿ä¸æ¬¡å¾ªç¯å®ææè½¬ã
*/
func rotate(s []int, r int) []int {
	lens := len(s)
	//åå»ºä¸ä¸ªç©ºçæå®é¿åº¦çslice
	res := make([]int, lens)
	for i := 0; i < lens; i++ {
		index := i + r
		if index >= lens {
			index = index - lens
		}
		res[i] = s[index]
	}
	return res
}

/*
ç»ä¹  4.5ï¼åä¸ä¸ªå½æ°å¨åå°å®ææ¶é¤[]stringä¸­ç¸é»éå¤çå­ç¬¦ä¸²çæä½ã
*/
func emptyString(s []string) []string {
	i := 0
	index := 0
	num := len(s)
	for _, v := range s {
		index = i + 1
		if index >= num {
			break
		}
		if v != s[index] {
			s[i] = v
			i++
		}
	}
	fmt.Println(s[:i])
	return s[:i]
}

/*
ç»ä¹  4.6ï¼ ç¼åä¸ä¸ªå½æ°ï¼åå°å°ä¸ä¸ªUTF-8ç¼ç ç[]byteç±»åçsliceä¸­ç¸é»çç©ºæ ¼ï¼åèunicode.IsSpaceï¼æ¿æ¢æä¸ä¸ªç©ºæ ¼è¿å
*/
func emptyString2(s []byte) []byte {
	index := 0
	num := len(s)
	for i := 0; i < num; i++ {
		index = i + 1
		num = len(s)
		if index >= num {
			break
		}
		if unicode.IsSpace(rune(s[i])) && unicode.IsSpace(rune(s[index])) {
			//ç»åremoveå½æ°
			copy(s[i:], s[index:])
			s = s[:len(s)-1]
			i--
		}
	}
	return s
}

/*
ç»ä¹  4.7ï¼ ä¿®æ¹reverseå½æ°ç¨äºåå°åè½¬UTF-8ç¼ç ç[]byteãæ¯å¦å¯ä»¥ä¸ç¨åéé¢å¤çåå­ï¼
*/
func reverse1(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
