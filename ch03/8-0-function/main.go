package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func factorials(n int) int {
	if n < 1 {
		return 1
	}
	return n * factorials(n-1)
}

func main() {
	// sum := add(10, 5)
	// fmt.Println(sum)
	// a, b := swap("foo", "bar")
	// fmt.Println(a, b)
	fmt.Println(factorials(11))
}
