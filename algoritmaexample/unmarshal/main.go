package main
 
import (
	"encoding/json"
	"io/ioutil"
)
 
type Salary struct {
        Basic, HRA, TA float64
    }
 
type Employee struct {
	FirstName, LastName, Email string
	Age                        int
	MonthlySalary             []Salary
}
 
func main() {
	data := Employee{
        FirstName: "Mark",
        LastName:  "Jones",
        Email:     "ismail@gmail.com",
        Age:       26,
        MonthlySalary: []Salary{
            Salary{
                Basic: 15000.00,
                HRA:   5000.00,
                TA:    2000.00,
            },
        },
    }
 
	file, _ := json.MarshalIndent(data, "", " ")
 
	_ = ioutil.WriteFile("test.json", file, 1)
}