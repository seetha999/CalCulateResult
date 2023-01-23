package main

import "fmt"

type Addnumber struct {
	value1 int
	value2 int
	value3 int
}

func main() {
	var a Addnumber
	a.value1 = 25
	a.value2 = 12
	a.value3 = a.value1 + a.value2
	fmt.Println("value is: ", a.value3)
}
