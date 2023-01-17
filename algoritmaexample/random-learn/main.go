package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for _, value := range rand.Perm(44) {
		fmt.Println(value)
	}
}
