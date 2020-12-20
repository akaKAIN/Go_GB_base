package models

import "fmt"

type LifeIndicator struct {
	title        string
	isVital      bool
	currentValue uint16
	maxValue     uint16
}

func GetNewLifeIndicator(title string, max uint16, isVital bool) *LifeIndicator {
	li := new(LifeIndicator)
	li.title, li.isVital, li.maxValue, li.currentValue = title, isVital, max, max
	return li
}

// Распечатать получить строкове представление структуры
func (li LifeIndicator) String() string {
	return fmt.Sprintf("%s: %v/%v", li.title, li.currentValue, li.maxValue)
}

// Получить текущее значение
func (li *LifeIndicator) Get() uint16 {
	return li.currentValue
}

// Отнимает значение от текущго значения. Возвращает логическое состояние после операции вычитания - жив или мертв
func (li *LifeIndicator) Loss(num uint16) (isDead bool) {
	if li.currentValue > num {
		li.currentValue -= num
		return false
	}
	// сюда попадаем только если переданное значение больше текущего.
	li.currentValue = 0
	if li.isVital {
		return true
	}
	return false
}

// Повысить масимальное значение и добавить разницу к текущему
func (li *LifeIndicator) UpMax(num uint16) bool {
	if num > li.maxValue {
		li.currentValue += num - li.maxValue
		li.maxValue = num
		return true
	}
	return false
}
