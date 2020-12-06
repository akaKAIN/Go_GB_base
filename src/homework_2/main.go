package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Результат: %d\n", makeExpression())
	fmt.Printf("Искомый массив нат.чисел: %v\n", GetSimpleNumbersCycle())
	fmt.Println(CalcExpression())
}
