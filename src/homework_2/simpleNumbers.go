package main

func GetSimpleArray(length int) []int {
	var arr = []int{2}
	for i := 3; i <= length; i += 2 {
		for ind, val := range arr {
			if i % val == 0 {
				break
			}
			if ind == len(arr) - 1 && i % val != 0 {
				arr = append(arr, i)
			}
		}
	}
	return arr
}

func GetSimpleNumbersCycle() []int {
	length := GetLengthFromInput()
	return GetSimpleArray(length)

}
