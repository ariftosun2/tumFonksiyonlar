 package main

import (
	"fmt"
	"sort"
)



func miniMaxSum(arr []int){
	maxSum,minSum,sum:=0,0,0
	sort.Slice(arr, func(a, b int) bool { return arr[a] < arr[b] })
	//sort.Ints(arr)
	fmt.Println(arr)
	for i:=0;i<len(arr);i++{
		sum+=arr[i]
		minSum=sum-arr[len(arr)-1]
		maxSum=sum-arr[0]
	}
	fmt.Println(minSum,maxSum)
}


func main(){
	rasgeleArr:=[]int{769082435, 210437958 ,673982045, 375809214, 380564127}
	miniMaxSum(rasgeleArr)
} 
/*package main

import (
	"fmt"
	"sort"
)

func calculator(arr []int, exe int) int {
	for i := 0; i < len(arr); i++ {
		exe += arr[i]
	}
	return exe
}

func miniMaxSum(arr []int) {
	maxSum, minSum := 0, 0
	a := len(arr)
	sort.Ints(arr)
	maxSum = calculator(arr[1:], maxSum)
	minSum = calculator(arr[:a-1], minSum)

	fmt.Println(minSum, maxSum)
}

func main() {
	rasgeleArr := []int{44, 55, 66, 22, 33, 56, 22}
	miniMaxSum(rasgeleArr)
} */