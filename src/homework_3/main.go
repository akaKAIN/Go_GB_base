package main

import (
	"fmt"
	"homework_3/bubble"
	"homework_3/fizzbuzz"
)

func main() {
	fmt.Println(bubble.Sort([]int{4, 3, 2, 1}))
	fizzbuzz.PrintFizzBuzz()
}
