package simplenum

import "github.com/akaKAIN/Go_GB_base/src/input"

/*
Сложность O(n log(log n))) взята из Wiki и с учетом того,
что проход осуществляется только по нечетным числам -
разделена на 2.
*/
func GetSimpleArray(length int) []int {
	var arr = []int{2}
	for i := 3; i <= length; i += 2 {
		for ind, val := range arr {
			if i%val == 0 {
				break
			}
			if ind == len(arr)-1 && i%val != 0 {
				arr = append(arr, i)
			}
		}
	}
	return arr
}

func GetSimpleNumbersCycle() []int {
	length := input.GetLengthFromInput()
	return GetSimpleArray(length)

}
