package bubble

import (
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
	tests := []struct {
		name     string
		array    []int
		expected []int
	}{
		{
			name:     "Random order",
			array:    []int{4, 5, 2, 7, 1},
			expected: []int{1, 2, 4, 5, 7},
		},
		{
			name:     "Already ordered",
			array:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Reverse order",
			array:    []int{5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := Sort(tc.array)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Fatalf("Expected: %v, gotted: %v\n", tc.expected, result)
			}
		})
	}
}

func BenchmarkSort(b *testing.B) {
	base := []int{5, 4, 3, 6, 9, 10, 7, 2, 1, 8}
	for i := 0; i < b.N; i++ {
		Sort(base)
	}
}
