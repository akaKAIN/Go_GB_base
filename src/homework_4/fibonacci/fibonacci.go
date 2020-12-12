package fibonacci

import (
	"math"
)


func calcFibonacciBenet(limit int) int {
	l := float64(limit)
	return int(math.Pow((1+math.Sqrt(5))/2, l-math.Pow((1-math.Sqrt(5))/2, l)/math.Sqrt(5)))
}

//func calcFibonacciRecursive(limit int) {
//	var result int
//	for i := 2; i < limit; i++ {
//
//	}
//}
