package main

import (
	"fmt"
	"homework_3/bubble"
	"homework_3/insertion"
)

func GetCopyArr(arr []int) []int {
	var newArr = make([]int, len(arr), len(arr))
	copy(newArr, arr)
	return newArr
}

func main() {
	fmt.Println(bubble.Sort(GetCopyArr([]int{4, 3, 2, 1})))
	fmt.Println(insertion.Sort(GetCopyArr([]int{4, 3, 2, 1})))
}
