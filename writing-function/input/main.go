package Input

import "fmt"

func Input() {
	fmt.Println("lutfen bir sayi girin:")
	var count int
	fmt.Scanln(&count)
	fmt.Println(count)
}
