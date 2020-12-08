package fizzbuzz

import "fmt"

func PrintFizzBuzz() {
	var out string
	for i := 1; i <= 100; i++ {
		out = ""
		if i%3 == 0 {
			out += "Fizz"
		}
		if i%5 == 0 {
			out += "Buzz"
		}
		if out == "" {
			fmt.Println(i)
		} else {
			fmt.Println(out)
		}
	}
}
