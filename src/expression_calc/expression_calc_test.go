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
			expectedError:   true,
			expression:      "234%12+3",
			expectedNums:    nil,
			expectedOperand: "",
		},
		{
			expectedError:   true,
			expression:      "234%12%3",
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

	for _, tc := range caseList {
		nums, operand, err := ParseOperands(tc.expression)
		errorExist := err != nil
		if errorExist != tc.expectedError {
			t.Fatalf("%s: Expected error: %v, but gotted: %v", tc.expression, tc.expectedError, err)
		}
		if !tc.expectedError {
			if !reflect.DeepEqual(nums, tc.expectedNums) {
				t.Fatalf("%s: Expected result nums: %q, but gotted: %q", tc.expression, tc.expectedNums, nums)
			}
			if operand != tc.expectedOperand {
				t.Fatalf("Expected operand: %q, but gotted: %q", tc.expectedOperand, operand)
			}
		}
	}
}
