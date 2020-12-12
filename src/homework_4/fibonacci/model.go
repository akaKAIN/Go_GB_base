package fibonacci

import (
	"fmt"
)

type FibNum struct {
	keys []int
	data map[int]int
}

// создаем базовое наполнение стрктуры
func New() *FibNum {
	var keys = []int{0, 1, 2}
	var baseVal = []int{0, 1, 1}
	var data = make(map[int]int)
	for i := 0; i < 3; i++ {
		data[keys[i]] = baseVal[i]
	}
	return &FibNum{
		keys: keys,
		data: data,
	}
}

func (f *FibNum) Get(key int) (int, error) {
	val, isExist := f.data[key]
	if isExist {
		return val, nil
	}
	return 0, fmt.Errorf("No keys mutches in data\n")

}

// Расчет ряда чисел Фибоначчи до указанного 'n'.
// Возвращает число Фибоначчи для указанного 'n'
func (f *FibNum) Calc(limit int) (int, error) {
	if limit < 0 {
		return 0, fmt.Errorf("Number must be positive\n")
	}
	for {
		fv, err := f.Get(limit)
		if err != nil {
			f.calcNext()
		} else {
			return fv, nil
		}
	}
}

// Расчитать значение Фибоначчи для последующего значения от имеющегося последнего в "мапе"
func (f *FibNum) calcNext() {
	newKey := f.keys[len(f.keys)-1] + 1             // Получаем значение последнего ключа и инкрементируем его на 1
	newValue := f.data[newKey-1] + f.data[newKey-2] // Получаем знач. Фибоначчи для нового ключа
	f.data[newKey] = newValue                       // Сохраняем ключ-значение в "мапу"
	f.keys = append(f.keys, newKey)                 // Добавляем новый ключ в массив с ключами
}
