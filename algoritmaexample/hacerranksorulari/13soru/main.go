package main

import "fmt"

func addTwoNumbers(l1 []int, l2 []int) []int {
	var b int
	var ListNode=  make([]int,len(l2))
	for i := 0; i < len(l1);i++{
		a := l1[i] + l2[i]
		b = 0
		if a >= 10 {
			b = a - 10
			ListNode = append(ListNode, b)
			l1[i+1]++
		}
		if a < 10 {
			
			ListNode = append(ListNode, a)

		}

	}
	return ListNode
}
func main() {
	zed := []int{1, 3, 4, 5, 6,9}
	yasua := []int{9, 3, 5, 6, 7,0}
	result := addTwoNumbers(zed, yasua)
	fmt.Println(result)
}
