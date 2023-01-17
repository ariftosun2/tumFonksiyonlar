package main

func myAtoi(s string) []int {
	var intgosterim []int
	var intgosterim2 []int
	
	for i,a=range s{
		intgosterim=append(intgosterim,i)
	} 
	for i,_=range intgosterim{
		if i=="-"{
			intgosterim2=append(intgosterim,[a])
		}
		if [a]<100000000000000000000000000000 {
			intgosterim2=append(intgosterim,[a])
		}
	}
	return intgosterim2
}


func main(){
	abc="llll3344433"
	abc2=myAtoi(abc)

}

