package fibonacci

import (
	"fmt"
)

type FibNum struct {
	num  int
	data map[int]int
}

// Создание базового наполнения стрктуры
// Возвращает базовую структуру с базовыми значениями для первых трех чисел ряда Фибоначчи
func New() *FibNum {
	var baseVal = []int{0, 1, 1}
	var data = make(map[int]int)
	var i int
	for ; i < 3; i++ {
		data[i] = baseVal[i]
	}
	return &FibNum{
		num:  i,
		data: data,
	}
}

// Метод получения значения Фибоначчи для "ключа" (n-нного числа в ряде Фибоначчи).
// Возвращает значение по ключу и ошибку, если ключа нет в "мапе".
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
	lastKey := f.num                                // Получаем последний имеющийся номер значения из ряда Фибоначчи
	newValue := f.data[lastKey-1] + f.data[lastKey] // Получаем знач. Фибоначчи для нового ключа
	lastKey += 1                                    // Увеличиваем номер рядя Ф. на 1 и производим сохранение значений:
	f.data[lastKey] = newValue                      // Сохраняем ключ-значение в "мапу"
	f.num = lastKey                                 // Сохраняем последний номер значения из ряда Ф.
}
