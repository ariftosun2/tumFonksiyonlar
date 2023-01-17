/* package main

import "fmt"



func fib(n int) int {
	if n == 2 {
		return 1
	} else if n == 1 {
		return 0
	}
	return fib(n-1) + fib(n-2)
}

func main() {
	a := fib(6)
	fmt.Println(a)
} */

package main

import "fmt"

func printFibonacciSeries(num int) {
	a := 0
	b := 1
	c := b
	fmt.Printf("Series is: %d %d", a, b)
	for i := 0; i <= num; i++ {
		c = b
		b = a + b
		if b >= num {
			fmt.Println()
			break
		}
		a = c
		fmt.Printf(" %d", b)
	}
}

func main() {
	printFibonacciSeries(7)
}
