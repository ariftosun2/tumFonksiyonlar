package main

import (
	"bufio"
	"fmt"
	"io"

	//"go/scanner"
	//"io"
	"log"
	"os"
	"strings"
)

type Vehicles struct {
	Name     string
	LastName string
}

func readerfile(filess string) *[]Vehicles {
	filec, err := os.Open(filess)
	if err != nil {
		log.Fatal(err)
	}
	defer filec.Close()
	scanners := bufio.NewScanner(filec)
	var render []Vehicles
	for scanners.Scan() {
		line := scanners.Text()
		record := strings.Split(line, ",")
		render = append(render, Vehicles{Name: record[0], LastName: record[1]})
	}
	return &render
}

func main() {
	lostfile, err := os.Create("testpacel.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer lostfile.Close()
	file1, err := os.OpenFile("testpacel.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file1.Close()

	d := []string{"Welcome to the world of Go1.", "Go is a compiled language.", "It is easy to learn Go."}
	arif:=createlines(file1,d)

	//filesss,_:=lostfile.WriteString("bakalim,buraya,eklenecekmi")
	vehislerss := readerfile(arif)
	fmt.Println(vehislerss)
}
func createlines(files string) []string{
	testdosyasi:=[]string{}

	for _,v:=range files{
		 	testdosyasi=append(testdosyasi,v )
	}
}
