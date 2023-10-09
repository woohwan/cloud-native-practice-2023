package main

import "fmt"

type Vertex struct {
	X, Y float64
}

func main() {
	// var v Vertex // 구조체는  nil이 될 수 없습니다.
	// fmt.Println(v)
	// v = Vertex{} // 명시적으로 빈 구조체 선언
	// fmt.Println(v)

	// v = Vertex{1.0, 2.0}
	// fmt.Println(v)
	// v = Vertex{Y: 2.5} // label을 이용하여 특정 필드 정의
	// fmt.Println(v)

	// 구초제 필드는 dot notation을 이용해 접근 가능
	// v := Vertex{X: 1.0, Y: 3.0}
	// fmt.Println(v) // {1 3}

	// v.X *= 1.5
	// v.Y *= 2.5
	// fmt.Println(v)

	var v *Vertex = &Vertex{1, 3}
	fmt.Println(v)
	v.X, v.Y = v.Y, v.X
	fmt.Println(v)

}
