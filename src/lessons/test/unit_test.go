package test

import (
	"testing"
)

func EvenNumber(a int) bool{
	return a%2 == 0
}

func Test_EvenNumber(t *testing.T) {
	sample := []int{1, 2, 3, 4, 5}
	expected := []bool{false, true, false, true, false}
	var actual []bool
	for _, v := range sample {
		actual = append(actual, EvenNumber(v))
	}
	for i := range actual {
		if actual[i] == expected[i] {
			continue
		} else {
			t.Error("index" , i, "in both array don't match")
		}
	}
	t.Log("Passed!")
}