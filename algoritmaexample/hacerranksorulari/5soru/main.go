package main

import (
	"fmt"
	"strings"
	//"strconv"
)

func staircase(n int) {
	a := " "
	b := "#"
	//dizedonusum,_:= strconv.Atoi(a)
	for i := 1; i <= n; i++ {

		fmt.Println(strings.Repeat(a, n-i) + strings.Repeat(b, i))
	}
}
func main() {

	staircase(99)
}
