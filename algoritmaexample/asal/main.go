package main

import (
	"fmt"
)

func asal(x int) bool {
	var b bool
	m := []int{2, 3, 4, 5, 6, 7, 8, 9}
	for i := 0; i < len(m); i++ {
		if x%m[0] == 1 {
			b = true
		}
	}
	return b
}
func main() {
	c := asal(1067)
	fmt.Println(c)
}
