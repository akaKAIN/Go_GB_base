package mathoperation

import "testing"

func Test_add(t *testing.T) {
	numA, numB, expectedResult := 3, 4, 7
	result := add(numA, numB)
	if result != expectedResult {
		t.Fatalf("Add error: expected %d, but got %d", expectedResult, result)
	}
}

func Test_sub(t *testing.T) {
	numA, numB, expectedResult := 3, 4, -1
	result := sub(numA, numB)
	if result != expectedResult {
		t.Fatalf("Sub error: expected %d, but got %d", expectedResult, result)
	}
}

func Test_mul(t *testing.T) {
	numA, numB, expectedResult := 4, 8, 32
	result := mul(numA, numB)
	if result != expectedResult {
		t.Fatalf("Mul error: expected %d, but got %d", expectedResult, result)
	}
}

func Test_div(t *testing.T) {
	numA, numB, expectedResult := 21, 3, 7
	result := div(numA, numB)
	if result != expectedResult {
		t.Fatalf("Div error: expected %d, but got %d", expectedResult, result)
	}

	numA, numB, expectedResult = 131, 0, 0
	result = div(numA, numB)
	if result != expectedResult {
		t.Fatalf("Div error: expected %d, but got %d", expectedResult, result)
	}
}
