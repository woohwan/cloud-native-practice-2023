package main

import (
	"fmt"
	"time"
)

func main() {
	// i := 0

	// switch i % 3 {
	// case 0:
	// 	fmt.Println("Zero")
	// 	fallthrough
	// case 1:
	// 	fmt.Println("One")
	// case 2:
	// 	fmt.Println("Two")
	// default:
	// 	fmt.Println("Huh?")
	// }

	hour := time.Now().Hour()

	switch {
	case hour >= 5 && hour < 9:
		fmt.Println("I'm writing")
	case hour > 9 && hour < 18:
		fmt.Println("I'm working")
	default:
		fmt.Println("I'm sleeping")
	}

}
