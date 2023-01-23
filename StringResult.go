package main

import "fmt"

type Name struct {
	name   string
	srname string
}

func (n Name) String() string {
	return fmt.Sprintf("My name is %q and sur name is %s.", n.name, n.srname)
}

func main() {
	n := Name{
		name:   "sita",
		srname: "Nannapaneni",
	}
	fmt.Println(n.String())
}
