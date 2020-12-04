package insertion

import (
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
	arr := []int{4, 5, 2, 7, 1}
	expectRes := []int{1, 2, 4, 5, 7}
	result := Sort(arr)
	if !reflect.DeepEqual(result, expectRes) {
		t.Fatalf("Expected: %v, gotted: %v\n", expectRes, result)
	}

	arr = []int{1, 2, 3, 4, 5}
	expectRes = []int{1, 2, 3, 4, 5}
	result = Sort(arr)
	if !reflect.DeepEqual(result, expectRes) {
		t.Fatalf("Expected: %v, gotted: %v\n", expectRes, result)
	}

	arr = []int{5, 4, 3, 2, 1}
	expectRes = []int{1, 2, 3, 4, 5}
	result = Sort(arr)
	if !reflect.DeepEqual(result, expectRes) {
		t.Fatalf("Expected: %v, gotted: %v\n", expectRes, result)
	}
}

func BenchmarkSort(b *testing.B) {
	base := []int{5, 4, 3, 6, 9, 10, 7, 2, 1, 8}
	expect := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < b.N; i++ {
		arr := Sort(base)
		if !reflect.DeepEqual(arr, expect) {
			b.Fatalf("Expected: %v, gotted: %v\n", expect, arr)
		}
	}
}
