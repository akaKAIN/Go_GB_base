package input

import (
	"bufio"
	"fmt"
	"github.com/akaKAIN/Go_GB_base/src/expression_calc"
	"github.com/akaKAIN/Go_GB_base/src/mathoperation"
	"io"
	"strconv"
	"strings"
)

/*
Функция содержащая полный цикл ввода и вычисления мат. выражения.
return: рузультат мат.вычислений.
*/
func makeExpression(reader io.Reader) int {
	var (
		numA, numB       int
		operationHandler func(int, int) int
	)
	numA =  getOperand(reader)
	operationHandler = getHandlerByOperator(reader)
	numB = getOperand(reader)

	return operationHandler(numA, numB)

}

/*
Запрос ввода от пользователя.
return: введеный текст
*/
func Input(msg string, reader io.Reader) string {
	var scanner = bufio.NewScanner(reader)
	fmt.Print(msg)
	scanner.Scan()
	text := scanner.Text()
	return text
}

/**
Парсинг числа введенного пользователем
*/
func getOperand(reader io.Reader) int {
	for {
		input := Input("Введите целое число: ", reader)
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
func getHandlerByOperator(reader io.Reader) func(int, int) int {
	operators := []string{"+", "-", "*", "/"}
	message := fmt.Sprintf("Введите один из операторов => (%s): ", strings.Join(operators, " "))
	operationsMap := mathoperation.GetOperationsMap()

	for {
		operator := Input(message, reader)
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
func GetLengthFromInput(reader io.Reader) int {
	for {
		input := Input("Введите верхний порог для поиска натурального числа: ", reader)
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
func CalcExpression(reader io.Reader) (result int) {
	operationsMap := mathoperation.GetOperationsMap()
	for {
		input := Input("Введите математическое выражение (пример: 256+32): ", reader)
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
