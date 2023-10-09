package main

import (
	"fmt"
	"os"
)

func main() {
	// defer fmt.Println("cruel world")

	// fmt.Println("goodbye")
	file, err := os.Create("./foo.txt")
	defer closeFile(file)
	if err != nil {
		return
	}

	_, err = fmt.Fprintln(file, "Your mother was a hamster")
	if err != nil {
		return
	}
	fmt.Println("File written to successfully")
}

func closeFile(f *os.File) {
	if err := f.Close(); err != nil {
		fmt.Println("Error closing file: ", err.Error())
	} else {
		fmt.Println("File closed successfully")
	}
}
