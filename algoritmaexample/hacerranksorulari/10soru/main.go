package main

import (
	"fmt"
	"sort"
	//"sort"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums1 = append(nums1, nums2...)
	sort.Ints(nums1)
	
	a := len(nums1)
	a2 := a % 2
	sum := 0.0

	if a2 == 0 {
		sum = float64((nums1[a/2] + nums1[a/2-1])) / 2

	} else {

		sum = float64(nums1[a/2])
	}
	fmt.Println(nums1)
	return sum
}
func main() {
	a1 := []int{22, 33, 44, 55}
	a2 := []int{22, 33, 44, 55}
	a3 := findMedianSortedArrays(a1, a2)
	fmt.Println(a3)
}
