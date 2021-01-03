package fizzbuzz

import (
	"fmt"
	"strconv"
)

func calcFizzBuzz(in int) string {
	var out string

	if in%3 == 0 {
		out += "Fizz"
	}
	if in%5 == 0 {
		out += "Buzz"
	}
	if out == "" {
		return strconv.Itoa(in)
	} else {
		return out
	}
}

func FizzBuzz() {
	for i := 0; i <= 100; i++ {
		fmt.Println(calcFizzBuzz(i))
	}
}
