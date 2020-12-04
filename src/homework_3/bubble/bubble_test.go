package bubble

import (
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
	arr := []int{4,5,2,7,1}
	expectRes := []int{1,2,4,5,7}
	result := Sort(arr)
	if !reflect.DeepEqual(result, expectRes) {
		t.Fatalf("Expected: %v, gotted: %v\n", expectRes, result)
	}

	arr = []int{1,2,3,4,5}
	expectRes = []int{1,2,3,4,5}
	result = Sort(arr)
	if !reflect.DeepEqual(result, expectRes) {
		t.Fatalf("Expected: %v, gotted: %v\n", expectRes, result)
	}

	arr = []int{5,4,3,2,1}
	expectRes = []int{1,2,3,4,5}
	result = Sort(arr)
	if !reflect.DeepEqual(result, expectRes) {
		t.Fatalf("Expected: %v, gotted: %v\n", expectRes, result)
	}
}
