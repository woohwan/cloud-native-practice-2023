package main

import "fmt"

func main() {
	// var a int = 10

	// var p *int = &a
	// fmt.Println(p)
	// fmt.Println(*p)

	var n *int
	var x, y int
	fmt.Println(n)
	fmt.Println(n == nil)

	fmt.Println(x == y)
	fmt.Println(&x == &x)
	fmt.Println(&x == &y)
	fmt.Println(&x == nil)

}
