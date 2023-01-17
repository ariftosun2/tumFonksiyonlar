package main

import "fmt"

func twoSum(nums []int, target int) []int {
	var sumReturn []int
	for i := 0; i < len(nums); i++ {
		if nums[i]+nums[i+1] == target {
			sumReturn = append(sumReturn, 0)
		}
		sumReturn = append(sumReturn, 1)
	}
	return sumReturn
}

func main() {
	var x[] int
	x = []int{1, 2, 3, 4, 5}

	fmt.Println(twoSum(x, 8))
}
