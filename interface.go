package main

import (
	"fmt"
)

type hedro interface {
	bunaSiera()
	halu(s string)
	increment()
	concrete() int
}
type hello struct {
	i int
}

func (tst *hello) bunaSiera() {
	fmt.Println("salutare")
}
func (tst *hello) halu(s string) {
	fmt.Println(s)
}
func (tst *hello) increment() {
	tst.i++
}
func (tst *hello) concrete() int {
	return tst.i

}
func main() {
	a := &hello{}
	a.bunaSiera()
	a.halu("hi there")
	a.increment()
	a.increment()
	fmt.Println(a.concrete())
}
