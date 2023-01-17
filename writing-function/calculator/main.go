package main

import "fmt"

type calculator interface {
	Multiply() float64
	Collects() float64
	Eject() float64
	Divide() float64
}
type values struct {
	A float64
	B float64
}

func (a *values) Multiply() float64 {
	multiply := a.A * a.B

	return multiply

}

func main() {
	carpim := Multiply()
	fmt.Println(carpim)

}
