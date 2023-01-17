package main

import (
	"fmt"
	"math"
)

func matrixsum(arr [][]int) int32 {
	sum1 := 0
	sum2 := 0
	size := len(arr)
	
	for i := 0; i < size; i++ {
		sum1 += arr[i][i]
		sum2 += + arr[i][size-1-i]
	}

	return int32(math.Abs(float64(sum1-sum2)))
}

func main() {
	test1 := [][]int{{1, 33, 3}, {4, 5, 6}, {7, 8, 20}}

	test2 := matrixsum(test1)
	fmt.Println(test2)
}
