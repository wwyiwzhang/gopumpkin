package gotest

func get_smallest_number(slice []int) int {
	smallest := slice[0]
	for _, val := range slice {
		if val < smallest {
			smallest = val
		}
	}
	return smallest
}
