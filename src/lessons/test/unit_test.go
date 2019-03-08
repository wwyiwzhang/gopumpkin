package unitest

import "testing"

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

func Benchmark_EvenNumber(b *testing.B) {
	for i:=0; i < b.N; i++ {
		EvenNumber(i)
	}
}