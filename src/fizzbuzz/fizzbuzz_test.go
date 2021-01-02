package fizzbuzz

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_calcFizzBuzz(t *testing.T) {
	tests := []struct {
		name string
		in   int
		out  string
	}{
		{name: "3", in: 3, out: "Fizz"},
		{name: "4", in: 4, out: "4"},
		{name: "5", in: 5, out: "Buzz"},
		{name: "6", in: 6, out: "Fizz"},
		{name: "7", in: 7, out: "7"},
		{name: "8", in: 8, out: "8"},
		{name: "9", in: 9, out: "Fizz"},
		{name: "15", in: 15, out: "FizzBuzz"},
		{name: "30", in: 30, out: "FizzBuzz"},
		{name: "31", in: 31, out: "31"},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.out, calcFizzBuzz(tc.in), "should be equal")
		})
	}
}
