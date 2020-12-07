package main

import (
	"bufio"
	"fmt"
	"homework_2/expression_calc"
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

/*
Получение от пользователя ввода с числом
верхней границей натуральный чисел.
*/
func GetLengthFromInput() int {
	for {
		input := Input("Введите верхний порог для поиска натурального числа: ")
		num, err := strconv.Atoi(input)
		if err != nil || num < 3 {
			fmt.Println("Некорректный ввод")
		} else {
			return num
		}
	}
}

/*
Получение от пользователя математического выражения
Возвращает результат вычисления
 */
func CalcExpression() (result int) {
	operationsMap := getOperationsMap()
	for {
		input := Input("Введите математическое выражение (пример: 256+32): ")
		nums, operator, err := expression_calc.ParseOperands(input)
		if err != nil {
			fmt.Println(err)
			continue
		} else {
			handlerFunc := operationsMap[operator]
			result = handlerFunc(nums[0], nums[1])
			break
		}
	}
	return
}
