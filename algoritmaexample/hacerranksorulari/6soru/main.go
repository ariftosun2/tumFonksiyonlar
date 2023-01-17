package main
import(
	"fmt"
	"strings"
)

func StairsAling(m  int){
	a:=" "
	b:="#"
	for i:=1;i<=m;i++{
		fmt.Println(strings.Repeat(a,m-i)+strings.Repeat(b,i)+strings.Repeat(b,i)+strings.Repeat(a,m-i))
	}

}

func main(){
	StairsAling(8)
}