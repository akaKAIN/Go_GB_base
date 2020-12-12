package main

import (
	"fmt"
	"homework_4/fibonacci"
)

func main() {
	n := 50
	fn := fibonacci.New()
	val, err := fn.Calc(n)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("struct", n, val)
	fmt.Println("recursion", n, fibonacci.CalcFibonacciRecursive(n))
	fmt.Println("bine", n, fibonacci.CalcFibonacciBine(n))
}
