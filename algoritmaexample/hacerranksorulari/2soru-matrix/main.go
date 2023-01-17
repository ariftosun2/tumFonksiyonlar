package main

import (
	"fmt"
	"math"
)

func matrixsum(arr [][]int) int {
	arrsum1:=0
   arrsum2:=0
	//var x,y int
		for i := 0; i < len(arr); i++ {
			for j := 0; j < len(arr); j++ {
				if i+j == len(arr)-1 {
					arrsum1 += arr[i][j]

				}
				if i == j {
					arrsum2 +=arr[i][j]
				}
			}
		
		}
		
	return int(math.Abs(float64(arrsum1)-float64(arrsum2)))

}

func main() {
	test1 := [][]int{{-99, 11, 9,10}, {3, 1, 6,10}, {4, 4, 9,10},{3,4,5,6}}


	test2 := matrixsum(test1)
	fmt.Println(test2)
}
