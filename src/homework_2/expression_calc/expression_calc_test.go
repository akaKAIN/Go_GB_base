package expression_calc

import (
	"reflect"
	"testing"
)

func TestParseOperands(t *testing.T) {
	expression := "10+20"
	expectedNums := []int{10, 20}
	expectedOperand := "+"
	nums, operand := ParseOperands(expression)
	if !reflect.DeepEqual(expectedNums, nums) {
		t.Fatalf("Expected nums: %d, but gotted: %d", expectedNums, nums)
	}
	if expectedOperand != operand {
		t.Fatalf("Expected operand: %s, but gotted: %s", expectedOperand, operand)
	}
}