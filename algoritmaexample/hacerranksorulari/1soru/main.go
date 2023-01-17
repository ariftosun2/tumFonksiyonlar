package main

import (
	"fmt"
	
)

func compareTriples(a ,b []int)[]int{
	alice:=0
	bob:=0
	var point []int
	for i:=0;i<len(a);i++{
		for c:=0;c<len(b);c++{
			if i==c {
				if a[i]>b[c]{
					alice+=1
				}else if a[i]<b[c]{
					bob+=1
				}
			}
		};
		
	}
	point=append(point,alice,bob)
	return point
}

func main(){
	a1:=[]int{10,20,30}
	a2:=[]int{40,50,30}
	toplam:=compareTriples(a1,a2)
	fmt.Println(toplam)

}