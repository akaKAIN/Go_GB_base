package expression_calc

import (
	"bytes"
	"fmt"
)

func ParseOperands(expression string) []int {
	exp := []byte(expression)
	if bytes.ContainsAny(exp, "+-*/") {
	}
	for ind, val := range exp {
		fmt.Printf("ind: %v, val: %v\n", ind, val)
	}
	return nil
}
