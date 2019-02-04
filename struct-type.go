package main

import (
	"fmt"
)

type person struct {
	Name   string
	Age    int
	Adress string
}

func main() {
	teo := person{
		Name:   "David T.",
		Age:    19,
		Adress: "Copenhagen,Denmark",
	}
	fmt.Println(teo)
	fmt.Println(teo.Name)
}

