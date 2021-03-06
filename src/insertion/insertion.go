package insertion

func Sort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		for j := i; j >= 1; j-- {
			if arr[j-1] > arr[j] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
			}
		}
	}
	return arr
}
