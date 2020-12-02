package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

func Input(msg string) string {
	var scanner = bufio.NewScanner(os.Stdin)
	fmt.Print(msg)
	scanner.Scan()
	text := scanner.Text()
	return text
}

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

func getHandlerByOperator() func(int, int) int {
	operators := []string{"+", "-", "*", "/"}
	message := fmt.Sprintf("Введите один из операторов => (%s): ", strings.Join(operators, " "))

	for {
		operator := Input(message)
		switch operator {
		// Тут конечно лучше было бы использовать словарь, но не успел с ним разобраться
		case "+":
			return add
		case "-":
			return sub
		case "*":
			return mul
		case "/":
			return div
		default:
			fmt.Printf("%q - не является доступным оператором\n", operator)
		}
	}
}

func GetLengthFromInput() int {
	for  {
		input := Input("Введите верхний порог для поиска натурального числа: ")
		num, err := strconv.Atoi(input)
		if err != nil || num < 3 {
			fmt.Println("Некорректный ввод")
		} else {
			return num
		}
	}
}
