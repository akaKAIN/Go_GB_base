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
		strArr    [2]string
		operators = "+-*/"
		intArr    = make([]int, 2, 2)
	)
	exp := []byte(expression)
	if !bytes.ContainsAny(exp, operators) {
		return nil, "", fmt.Errorf("Неверное выражение - неизвестный математический оператор.\n")
	}

	for _, o := range operators {
		// Пропускаем первый символ выражения (если число отрицательное)
		if operandInd := strings.LastIndex(expression, string(o)); operandInd != -1 && operandInd != 0 {
			if operandInd >= len(exp) {
				return nil, "", fmt.Errorf("Позиция оператора - неверная\n")
			}
			operand = string(expression[operandInd])
			strArr[0], strArr[1] = string(exp[:operandInd]), string(exp[operandInd+1:])
		}
	}

	for ind, val := range strArr {
		if num, err := strconv.Atoi(val); err != nil {
			return nil, "", fmt.Errorf("Неверное выражение - введены некорректные числа\n")
		} else {
			intArr[ind] = num
		}
	}

	return intArr, operand, nil
}
