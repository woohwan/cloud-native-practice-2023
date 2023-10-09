package main

import "fmt"

func main() {
	s := "foo"
	r := []rune(s) // casting
	b := []byte(s)

	fmt.Printf("%7T %v\n", s, s)
	fmt.Printf("%7T %v\n", r, r)
	fmt.Printf("%7T %v\n", b, b)

}
