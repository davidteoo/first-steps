package main

import (
	"fmt"
)

func main() {
	//defer fmt.Println("you") //defer==>it will run the command after the program finish
	//defer fmt.Println("meet")// <==this will be the first "after-command"
	fmt.Println("nice to")
	testpanics()
	fmt.Println("meet")
}
func testpanics() {
	defer func() { //it will recover after ERROR404
		if err := recover(); err != nil { //nil=zero value for pointers, interfaces, maps, slices, channels and function types
			fmt.Println(err)
			fmt.Println("the error has been solved")
		}
	}()
	panic("ERROR404")
}
