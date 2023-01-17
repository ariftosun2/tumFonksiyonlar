package main

import (
	"fmt"
	"sort"
	//"sort"
)

func resevre(str string) string {
	byte_str := []rune(str)
	for i, j := 0, len(byte_str)-1; i < j; i, j = i+1, j-1 {
		byte_str[i], byte_str[j] = byte_str[j], byte_str[i]
	}
	return string(byte_str)

}

func lengthComparison(sliceLegth []string) {
	var ListString []int

	m := make(map[string]int)
	
	for i := 0; i < len(sliceLegth); i++ {
		m[sliceLegth[i]] = len(sliceLegth[i])
	}
	/* 	for i,v:=range m{
		//ListString=append(ListString, v)

		result:=len(m)
	} */
	for _, b := range m {
		ListString = append(ListString, b)

	}
	sort.Ints(ListString)

	for i, v := range m {
		stringlen := len(ListString) - 1
		a1 := ListString[stringlen]

		if a1 == v {
			fmt.Println(i)
		}
	}

}

func longestPalindrome(s string) {
	var ListNode []string
	for i := 0; i < len(s); i++ {
		for a := i + 1; a < len(s); a++ {
			if s[i] == s[a] {
				ss := s[i : a+1]
				if ss == resevre(ss) {
					ListNode = append(ListNode, ss)
				}
			}
		}
	}
	lengthComparison(ListNode)
}

func main() {
	names := "llaa"
	longestPalindrome(names)
	
}
