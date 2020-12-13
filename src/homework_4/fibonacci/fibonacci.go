package fibonacci

import (
	"math"
)

// Получение числа Фибоначчи по формуле Бине. Никаких рекурсий и итераций.
func CalcFibonacciBine(limit int) int {
	l := float64(limit)
	p1 := (1 + math.Sqrt(5)) / 2
	p2 := (1 - math.Sqrt(5)) / 2
	p := math.Pow(p1, l) - math.Pow(p2, l)
	return int(math.Round(p / math.Sqrt(5)))
}

// Получение Фибоначи через рекурсию
func CalcFibonacciRecursive(limit int) int {
	if limit == 0 || limit == 1 {
		return limit
	}
	return CalcFibonacciRecursive(limit-2) + CalcFibonacciRecursive(limit-1)
}
