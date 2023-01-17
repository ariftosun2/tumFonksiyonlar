package main

import (
	"fmt"
	"strconv"
)

func reverse(x int) int {
	a:=strconv.Itoa(x)
	l:=len(a)
	eksi:="-"
	
	
	var reverseBox []string
	for i:=l;i>=0;i--{
		if a[i]==(eksi){
			 reverseBox=append(reverseBox, "-")
		}else{
		reverseBox=append(reverseBox,string(a[i]))
		}
	}
	donus,_:=strconv.Atoi(reverseBox[0])
	return donus
}


func main(){
	arif:=463
	fmt.Println(reverse(arif))
}