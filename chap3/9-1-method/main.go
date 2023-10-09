package main

import "fmt"

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Square() { // method를 *Vertex 타입에 부탁
	v.X *= v.X
	v.Y *= v.Y
}

type MyMap map[string]int

func (m MyMap) Length() int {
	return len(m)
}

func main() {
	// vert := &Vertex{3, 4}
	// fmt.Println(vert)

	// vert.Square()
	// fmt.Println(vert)
	mm := MyMap{"A": 1, "B": 3}
	fmt.Println(mm)
	fmt.Println(mm["A"])
	fmt.Println(mm.Length())
}
