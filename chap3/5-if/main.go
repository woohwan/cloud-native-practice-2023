package main

import (
	"fmt"
	"os"
)

func main() {
	// if 7%2 == 0 {
	// 	fmt.Println("7 is even")
	// } else {
	// 	fmt.Println("7 is odd")
	// }

	// if 구문내에서 조건절을 수행하기 위한 초기화 구문 사용가능
	if _, err := os.Open("foo.ext"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("All is file")
	}
}
