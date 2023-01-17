package main

import "fmt"

func packaging(count int, consecutive_count int, start_count int) map[int][]int {

	package_count := make(map[int][]int)

	for i := 0; i <= count; i++ {
		thisNote := map[int]int{
			i: start_count,
		}

		package_count[i] = append(package_count[i], thisNote[i])
		start_count += consecutive_count

	}

	return package_count
}

func main() {

	var packages = make(map[int][]int)

	packages = packaging(56, 5, 7)
	fmt.Println(packages)

}
