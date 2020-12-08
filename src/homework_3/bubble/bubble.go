package bubble

func Sort(arr []int) []int {
	// цикл пузыльковой сортировки
	for {
		done := true
		for i := 0; i < len(arr)-1; i++ {
			if arr[i] > arr[i+1] {
				done = false
				arr[i], arr[i+1] = arr[i+1], arr[i]
			}
		}
		if done {
			break
		}
	}

	return arr
}
