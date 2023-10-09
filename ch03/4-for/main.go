package main

import "fmt"

func main() {
	// sum := 0
	// for i := 0; i < 10; i++ {
	// 	sum += 1
	// }

	// sum, i := 0, 0
	// for i < 10 {
	// 	sum += i
	// 	i++
	// }

	// fmt.Println(i, sum)

	// fmt.Println("For ever")
	// for {
	// 	fmt.Println("...and ever")
	// }

	// for & range
	// s := []int{2, 4, 8, 16, 32}
	// for i, v := range s {
	// 	fmt.Println(i, "->", v)
	// }

	// 불필요한 값: _
	// a := []int{0, 2, 4, 6, 8}
	// sum := 0

	// for _, v := range a {
	// 	sum += v
	// }
	// fmt.Println(sum)

	// 맵에 대한 반복문
	m := map[int]string{
		1: "Januaary",
		2: "Feburary",
		3: "March",
		4: "April",
	}
	for k, v := range m {
		fmt.Println(k, "->", v)
	}
}
