package main

import "fmt"

func main() {
	num := 5
	f := plusOne(num)
	fmt.Println(f())
	fmt.Println(f())
	division(5, 0)
}

// closure
func plusOne(x int) func() int {
	return func() int {
		x++
		return x
	}
}

// defer (error handling)
func division(x, y int) int {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("The denominator cannot be 0")
		}
	}()
	return x / y
}
