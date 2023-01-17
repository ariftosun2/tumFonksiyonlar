package main

import (
	"fmt"
	"math/rand"
)

func throwPebble(r float64) func() bool {

	return func() bool {
		x := rand.Float64()*2 - r
		y := rand.Float64()*2 - r
		if (x*x + y*y) <= r {
			return true
		} else {
			return false
		}
	}

}

func simulatePi(n int) {
	r := 1.0
	sumCircle := 0
	s := throwPebble(r)
	for i := 0; i < n; i++ {
		if s() {
			sumCircle++
		}
	}
	piResult := 4.0 * float64(sumCircle) / float64(n)

	fmt.Println(piResult)
}

func main() {
	simulatePi(100000)
}
