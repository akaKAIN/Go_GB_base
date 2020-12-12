package fibonacci

import "testing"

type CaseFibonacci struct {
	n        int
	expected int
}

var CaseList = []CaseFibonacci{
	{
		n:        0,
		expected: 0,
	},
	{
		n:        2,
		expected: 1,
	},
	{
		n:        3,
		expected: 2,
	},
	{
		n:        8,
		expected: 21,
	},
	{
		n:        10,
		expected: 55,
	},
	{
		n:        19,
		expected: 4181,
	},
}

func Test_fibonacciBinet(t *testing.T) {
	for _, testCase := range CaseList {
		f := calcFibonacciBenet(testCase.n)
		if f != testCase.expected {
			t.Fatalf("Expected %d, but gotted %d", testCase.expected, f)
		}
	}
}
