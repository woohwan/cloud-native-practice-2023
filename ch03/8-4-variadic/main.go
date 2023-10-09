package main

import "fmt"

func main() {
	// const name, age = "Kim", 22
	// fmt.Printf("%s is %d years old.\n", name, age)
	fmt.Println(product(2, 2, 2))

	m := []int{3, 3, 3}
	fmt.Println(product(m...))
}

func product(factors ...int) int {
	p := 1
	for _, n := range factors {
		p *= n
	}
	return p
}
