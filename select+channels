package main

import (
	"fmt"
	"time"
)

func main() {
	select {
	case f1 := <-onemoment(3, 2):
		fmt.Println(f1) //will select the fastest value,in this case f2
	case f2 := <-onemoment(10, 1):
		fmt.Println(f2)
		/*  default:
		    fmt.Println ("all channels are slow")*/
	}
}
func onemoment(v, i int) chan int { //will wait for i seconds before sending value v on the return channel
	rch := make(chan int)
	go func() {
		time.Sleep(time.Duration(i) * time.Second)
		rch <- v
	}()
	return rch
}
