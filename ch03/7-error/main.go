package main

import (
	"fmt"
)

func main() {

	// file, err := os.Open("somefile.ext")

	// if err != nil {
	// 	log.Fatal(err)
	// 	return err
	// }

	// e1 := errors.New("error 42")
	// e2 := fmt.Errorf("error %d", 42)

}

type NestedError struct {
	Message string
	Err     error
}

func (e NestedError) Error() string {
	return fmt.Sprintf("%s\n contains: %s", e.Message, e.Err.Error())
}
