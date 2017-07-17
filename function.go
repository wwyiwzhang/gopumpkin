package main

import "fmt"

func main() {
	newSlice := []int{10, 15, 20}
	smallest_value := get_smallest_number(newSlice)
	fmt.Println("The smallest number is", smallest_value)
	largest_value := get_largest_number(newSlice)
	fmt.Println("The largest number is", largest_value)
	first, second := get_first_second_element(newSlice)
	fmt.Println("The first element is", first)
	fmt.Println("The second element is", second)
	fmt.Println("The sum is", add(1, 2, 3, 4, 5))
}

// return single value
func get_smallest_number(slice []int) int {
	smallest := slice[0]
	for _, val := range slice {
		if val < smallest {
			smallest = val
		}
	}
	return smallest
}

func get_largest_number(slice []int) int {
	largest := slice[0]
	for _, val := range slice {
		if val > largest {
			largest = val
		}
	}
	return largest
}

// return multiple values
func get_first_second_element(slice []int) (int, int) {
	return slice[0], slice[1]
}

// unlimited parameters
func add(args ...int) int {
	sum := 0
	for _, val := range args {
		sum = sum + val
	}
	return sum
}
