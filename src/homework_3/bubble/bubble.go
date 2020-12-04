package bubble

func Sort(arr []int) []int {
	// Копируем слайс
	var sortedArr = make([]int, len(arr))
	copy(sortedArr, arr)

	// цикл пузыльковой сортировки
	for {
		done := true
		for i := 0; i <= len(sortedArr)-2; i++ {
			if sortedArr[i] > sortedArr[i+1] {
				done = false
				sortedArr[i], sortedArr[i+1] = sortedArr[i+1], sortedArr[i]
			}
		}
		if done {
			break
		}
	}

	return sortedArr
}
