package main

import (
	"fmt"
)

func main() {
	s0 := []int{0, 1, 2, 3, 4, 5, 6} // slice literal
	fmt.Println(s0)

	s1 := s0[:4]
	fmt.Println(s1)

	s2 := s0[3:]
	fmt.Println(s2)

	s0[3] = 42
	fmt.Println(s0)
	fmt.Println(s1)
	fmt.Println(s2)

}
