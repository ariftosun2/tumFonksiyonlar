package main

import "fmt"

//
func Intersection(firtstArray []string, secondArray []string) []string {
	var tanimliKume []string
	for _, a := range firtstArray {
		for _, b := range secondArray {
			if a == b {
				tanimliKume = append(tanimliKume, a)
			}

		}

	}
	return tanimliKume
}
func main() {
	firtstArray := []string{"arif", "tosun", "ismail", "yusuf"}
	secondArray := []string{"arif", "ismail", "yusuf", "abdulkadir", "eyup"}
	benzeyenler := Intersection(firtstArray, secondArray)
	fmt.Println(benzeyenler)
}
