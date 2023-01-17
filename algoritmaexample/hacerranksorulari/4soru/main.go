package main

import "fmt"


func pulusMinus(arr []int) {
	var sumzero float64
	var sumpositif float64
	var sumnegatif float64
	sumzero = 0
	sumpositif = 0
	sumnegatif = 0
	for i := 0; i < len(arr); i++ {
		if arr[i] < 0 {
			sumnegatif += 1

		} else if arr[i] > 0 {
			sumpositif += 1

		} else {
			sumzero += 1
		}

	}
	fmt.Println(sumpositif/float64(len(arr)))
	fmt.Println(sumzero/float64(len(arr)))
	fmt.Println(sumnegatif / float64(len(arr)))
}

func main() {
	arrTest := []int{0,0,1,-1,1}
	 pulusMinus(arrTest)
	

}
