package main

import "fmt"

func main() {
	// freezing := make(map[string]float32)

	// freezing["celcius"] = 0.0
	// freezing["fahrenheit"] = 32.0
	// freezing["kelvin"] = 273.2

	// fmt.Println(freezing["kelvin"])
	// fmt.Println(len(freezing))

	// delete(freezing, "kelvin")
	// fmt.Println(freezing)

	// map 초기화
	freezing := map[string]float32{
		"celsisus":   0.0,
		"fahrenheit": 32.0,
		"kelvin":     273.2,
	}
	fmt.Println(freezing)

	foo := freezing["no-such-key"]
	fmt.Println(foo)

	newton, ok := freezing["newton"]
	fmt.Println(newton)
	fmt.Println(ok)
}
