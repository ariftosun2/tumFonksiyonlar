package main
import "fmt"

func FirstFactorial(num int) int {

  
  a:=1
  if num==0{
	  return 1
  }
  for i:=1;i<=num;i++{
    a*=i
  }
  
  return a;

}

func main() {

  fmt.Println(FirstFactorial(0))

}