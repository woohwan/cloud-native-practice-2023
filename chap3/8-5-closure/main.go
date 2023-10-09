package main

import "fmt"

func sum(x, y int) int {
	return x + y
}
func product(x, y int) int {
	return x * y
}

func main() {
	var f func(int, int) int // 함수로 선언된 변수는 타입을 갖습니다.

	f = sum
	fmt.Println(f(3, 5))

	f = product
	fmt.Println(f(3, 5))
}
