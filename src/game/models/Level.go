package models

var (
	baseLevel       uint8  = 1
	baseCurrentExp  uint64 = 0
	expForNextLevel uint64 = 100
)

// Структура с информацией о текущем уровне
type Level struct {
	currentLevel      *uint8
	currentExperience *uint64
	expForNextLevel   *uint64
}

// Получение стукруры с информацией о текущем уровне
func GetStartLevel() *Level {
	return &Level{
		currentLevel:      &baseLevel,
		currentExperience: &baseCurrentExp,
		expForNextLevel:   &expForNextLevel,
	}
}

//увеличение счетчика уровня и порога опыта до следующего уровня
func (l *Level) LevelIncrement() {
	if *l.currentExperience > *l.expForNextLevel {
		*l.currentLevel++
		rate := uint64(*l.currentLevel)
		*l.expForNextLevel += rate * expForNextLevel
	}
}

// Получение информации о текущем уровне
func (l Level) GetCurrentLevel() uint8 {
	return *l.currentLevel
}
