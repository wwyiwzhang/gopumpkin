package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func handleError(err error) {
	if err != nil {
		fmt.Print("File Rreading error!")
		fmt.Print(err)
	}
}

func main() {
	// read file
	file_a, err := ioutil.ReadFile("some file path")
	handleError(err)
	fmt.Printf("%s\n", string(file_a))

	file_b, err := os.Open("some file path")
	handleError(err)
	b1 := make([]byte, 5)
	data, err := file_b.Read(b1)
	handleError(err)
	fmt.Printf("%d bytes are:", data, string(b1))
	file_b.Close()

	// write file
}
