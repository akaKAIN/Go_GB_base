package expression_calc

import (
	"reflect"
	"testing"
)

func TestParseOperands(t *testing.T) {
	var (
		expectedError error
		expression = "10+20"
		expectedNums = []int{10, 20}
		expectedOperand = "+"
	)
	nums, operand, err := ParseOperands(expression)
	if err != nil {
		t.Fatalf("Expected error: %q, but gotted: %q", expectedError, err)
	}
	if !reflect.DeepEqual(expectedNums, nums) {
		t.Fatalf("Expected nums: %d, but gotted: %d", expectedNums, nums)
	}
	if expectedOperand != operand {
		t.Fatalf("Expected operand: %s, but gotted: %s", expectedOperand, operand)
	}

	expression = "df+23"
	nums, operand, err = ParseOperands(expression)
	if err == nil {
		t.Fatalf("Expected error, but gotted: %q", err)
	}

	expression = "23-122+12"
	nums, operand, err = ParseOperands(expression)
	if err == nil {
		t.Fatalf("Expected error, but gotted: %q", err)
	}

	expression = "233"
	nums, operand, err = ParseOperands(expression)
	if err == nil {
		t.Fatalf("Expected error, but gotted: %q", err)
	}

	expression = "234%123"
	nums, operand, err = ParseOperands(expression)
	if err == nil {
		t.Fatalf("Expected error, but gotted: %q", err)
	}

	expression = "-111+11"
	expectedNums = []int{-111, 11}
	expectedOperand = "+"
	nums, operand, err = ParseOperands(expression)
	if err == nil {
		t.Fatalf("Expected error, but gotted: %q", err)
	}
}