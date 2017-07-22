package gotest

import "testing"

func Test_get_smallest_number(t *testing.T) {
	testArray := []int{1, 4, 5, 6, 9, 0, 3}
	smallest := get_smallest_number(testArray)
	if smallest == 0 {
		t.Log("Test Passed!")
	} else {
		t.Error("Test Failed!")
	}
}
