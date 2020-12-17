package simplenum

import (
	"reflect"
	"testing"
)

func TestGetSimpleArray(t *testing.T) {
	var (
		length                      int
		expectedResultArray, result []int
	)

	length = 37
	expectedResultArray = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37}
	result = GetSimpleArray(length)
	if !reflect.DeepEqual(result, expectedResultArray) {
		t.Fatalf("Expected array: %v\ngotted array: %v\n", expectedResultArray, result)
	}

	length = 3
	expectedResultArray = []int{2, 3}
	result = GetSimpleArray(length)
	if !reflect.DeepEqual(result, expectedResultArray) {
		t.Fatalf("Expected array: %v\ngotted array: %v\n", expectedResultArray, result)
	}
}
