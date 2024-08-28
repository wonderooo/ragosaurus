package util

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	result := Merge([]int{1, 2}, []int{3, 4})
	expected := []int{1, 2, 3, 4}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Content of slice not correct: %d, expected: %d", result, expected)
	}
}
