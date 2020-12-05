package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
Функция содержащая полный цикл ввода и вычисления мат. выражения.
return: рузультат мат.вычислений.
*/
func makeExpression() int {
	var (
		numA, numB       int
		operationHandler func(int, int) int
	)
	numA = getOperand()
	operationHandler = getHandlerByOperator()
	numB = getOperand()

	return operationHandler(numA, numB)

}

/*
Запрос ввода от пользователя.
return: введеный текст
*/
func Input(msg string) string {
	var scanner = bufio.NewScanner(os.Stdin)
	fmt.Print(msg)
	scanner.Scan()
	text := scanner.Text()
	return text
}

/**
Парсинг числа введенного пользователем
 */
func getOperand() int {
	for {
		input := Input("Введите целое число: ")
		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Printf("Значение %s не является целым числом.\n", input)
		} else {
			return num
		}
	}
}

/*
Получение от пользователя ввода математического оператора
return: функцию соответствующую введеному мат.оператору
 */
func getHandlerByOperator() func(int, int) int {
	operators := []string{"+", "-", "*", "/"}
	message := fmt.Sprintf("Введите один из операторов => (%s): ", strings.Join(operators, " "))
	operationsMap := getOperationsMap()

	for {
		operator := Input(message)
		handlerFunc, isExists := operationsMap[operator]
		if isExists {
			return handlerFunc
		} else {
			fmt.Printf("%q - не является доступным оператором\n", operator)
		}
	}
}
