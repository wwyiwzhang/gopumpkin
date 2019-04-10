package main

import (
	"reflect"
	"sync"
	"testing"
)

func Test_Count(t *testing.T) {
	wg := &sync.WaitGroup{}
	testCases := []struct {
		userList []string
		times    int
		expected map[string]int
	}{
		{
			userList: []string{"minion1", "minion2", "minion3"},
			times:    5,
			expected: map[string]int{
				"minion1": 5,
				"minion2": 5,
				"minion3": 5,
			},
		},
		{
			userList: []string{"elsa", "anna", "mulan"},
			times:    3,
			expected: map[string]int{
				"elsa":  3,
				"anna":  3,
				"mulan": 3,
			},
		},
	}
	for _, tc := range testCases {
		newCounter := NewUserCounter()
		actual := newCounter.Count(wg, tc.userList, tc.times)
		if !reflect.DeepEqual(actual, tc.expected) {
			t.Errorf("expected: %v, got: %v", tc.expected, actual)
		}
	}
}
