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
	newKey := f.num + 1                             // Получаем значение последнего ключа и инкрементируем его на 1
	newValue := f.data[newKey-1] + f.data[newKey-2] // Получаем знач. Фибоначчи для нового ключа
	f.data[newKey] = newValue                       // Сохраняем ключ-значение в "мапу"
	f.num = newKey                                  // Добавляем новый ключ в массив с ключами
}
