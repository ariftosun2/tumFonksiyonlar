package main

import (
	"fmt"
	"math"
)

func addTwoNumbers(l1, l2 []int) []int {
	var ListNode []int
	a := len(l1)
	b := len(l2)
	ss := math.Abs(float64(a - b))
	for s := 0; s < int(ss); s++ {
		if a > b {
			l2 = append(l2, 0)
		} else if a < b {
			l1 = append(l1, 0)
		} else {
			fmt.Println(l1, l2)
		}
	}
	h := len(l1)
	for i := 0; i < h; i++ {
		if l1[i]+l2[i] >= 10 {
			ListNode = append(ListNode, (l1[i] + l2[i] - 10))
			l1[i+1] += 1
		} else {
			ListNode = append(ListNode, (l1[i] + l2[i]))
		}
	}
	return ListNode
}
func main() {
	a1 := []int{2, 3, 4, 6, 1, 5, 9, 8, 5}
	a2 := []int{3, 4, 5, 6, 1, 5, 4, 8, 6, 2, 1}
	result := addTwoNumbers(a1, a2)
	fmt.Println(result)
}

