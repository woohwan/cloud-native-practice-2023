package main

import "fmt"

func incrementer() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	increment := incrementer()
	fmt.Println(increment()) // 1
	fmt.Println(increment())
	fmt.Println(increment())

	newIncrement := incrementer()
	fmt.Println(newIncrement())

}
