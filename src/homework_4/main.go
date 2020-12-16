package main

import (
	"bufio"
	"fmt"
	"homework_4/fibonacci"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n int64
	for {
		fmt.Print("Введите положительный номер N-ого числа из ряда Фибоначчи: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			continue
		}
		input = strings.TrimSpace(input)
		n, err = strconv.ParseInt(input, 10, 64)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if n < 0 {
			fmt.Println("Число должно быть положительным.")
			continue
		} else {
			break
		}
	}
	N := int(n)
	fn := fibonacci.New()
	val, err := fn.Calc(N)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("struct", n, val)
	fmt.Println("recursion", n, fibonacci.CalcFibonacciRecursive(N))
	fmt.Println("bine", n, fibonacci.CalcFibonacciBine(N))
}
