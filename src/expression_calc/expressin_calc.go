package expression_calc

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

func ParseOperands(expression string) ([]int, string, error) {
	var (
		operand   string
		operators = "+-*/"
		strArr    []string
		intArr    = make([]int, 2, 2)
	)
	exp := []byte(expression)
	if !bytes.ContainsAny(exp, operators) {
		return nil, "", fmt.Errorf("Введено неверное выражение.\n")
	}
	for _, val := range exp {
		if strings.ContainsRune(operators, rune(val)) {
			operand = string(val)
			break
		}
	}
	strArr = strings.Split(expression, operand)
	if len(strArr) == 2 {

		// Приводим строковоые значения к числу и добаявляем в слайс с числами.
		for ind, val := range strArr {
			if num, err := strconv.Atoi(val); err != nil {
				return nil, "", fmt.Errorf("Неверное выражение - введены ннкорректные числа\n")
			} else {
				intArr[ind] = num
			}
		}
	} else {
		return nil, "", fmt.Errorf("Неверное выражение - введено слишком много операторов.\n")
	}

	return intArr, operand, nil
}
