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
	//read file
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
	numbers := []byte("1\n2\n3\n4\n5")
	file_c, err := os.Create("some file path")
	handleError(err)

	byte_info, err := file_c.Write(numbers)
	handleError(err)
	fmt.Printf("number of byte written: %d\n", byte_info)

	file_c.Close()

	file_d, err := ioutil.ReadFile("some file path")
	handleError(err)

	fmt.Println("data is:\n", string(file_d))

}
