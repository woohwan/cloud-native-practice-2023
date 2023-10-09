package main

import "fmt"

func main() {
	// x := 5

	// zeroByValue(x)
	// fmt.Println(x)

	// ZeroByReference(&x)
	// fmt.Println(x)
	m := map[string]int{"a": 0, "b": 1}
	fmt.Println(m)

	update(m)
	fmt.Println(m)
}

func zeroByValue(x int) {
	x = 0
}

func ZeroByReference(x *int) {
	*x = 0
}

func update(m map[string]int) {
	m["c"] = 2
}
