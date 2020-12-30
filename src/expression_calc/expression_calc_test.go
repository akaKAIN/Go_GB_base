package expression_calc

import (
	"reflect"
	"testing"
)

type caseTest struct {
	expectedError   bool
	expression      string
	expectedNums    []int
	expectedOperand string
}

func TestParseOperands(t *testing.T) {
	caseList := []caseTest{
		{
			expectedError:   false,
			expression:      "10+20",
			expectedNums:    []int{10, 20},
			expectedOperand: "+",
		},
		{
			expectedError:   true,
			expression:      "df+23",
			expectedNums:    nil,
			expectedOperand: "",
		},
		{
			expectedError:   true,
			expression:      "23-122+12",
			expectedNums:    nil,
			expectedOperand: "",
		},
		{
			expectedError:   true,
			expression:      "233",
			expectedNums:    nil,
			expectedOperand: "",
		},
		{
			expectedError:   true,
			expression:      "234%123",
			expectedNums:    nil,
			expectedOperand: "",
		},
		{
			expectedError:   false,
			expression:      "-111+11",
			expectedNums:    []int{-111, 11},
			expectedOperand: "+",
		},
	}

	for _, test := range caseList {
		nums, operand, err := ParseOperands(test.expression)
		if (err == nil) == test.expectedError {
			t.Fatalf("Expected error: %v, but gotted: %v", test.expectedError, err)
		}
		if !test.expectedError {
			if !reflect.DeepEqual(nums, test.expectedNums) {
				t.Fatalf("Expected result nums: %q, but gotted: %q", test.expectedNums, nums)
			}
			if operand != test.expectedOperand {
				t.Fatalf("Expected operand: %q, but gotted: %q", test.expectedOperand, operand)
			}
		}
	}
}
