package main

import (
	"fmt"
)

func main() {
	//a := [5]int {1,2,3,4,5}
	myslice := []int{2, 3, 44}
	myslice = append(myslice, 71, 21, 2) // introduce values
	fmt.Println(myslice)
	s := make([]int, 5)
	s[0], s[1], s[2], s[3], s[4] = 1, 2, 3, 6, 7
	fmt.Println(s)
	s1 := s[1:3]

	fmt.Println(s1[:cap(s1)]) //capacity
	s2 := make([]int, 6)
	copy(s2, s[0:3]) // copy 1,2,3,6 from s[] to s[0],s[1]...

	fmt.Println(s2)
	fmt.Println(s2[:len(s2)]) //lenght

	copy(s2, s[0:5])
	fmt.Println(s2)

	/*x:= make (map[string]int)  //2 ways to declare maps
	  x["first"]=1
	  x["second"]=2
	*/
	x := map[string]int{
		"first":  1,
		"second": 2,
	}

	fmt.Println(x["first"])
	if v, ok := x["second"]; ok {
		fmt.Println(v)
	}
	fmt.Println(x)
	delete(x, "first") //delete a sequence from a map
	fmt.Println(x)
}
