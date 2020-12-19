package models

import "fmt"

type Person struct {
	Nickname  string
	Level     *Level
	Class     *GameClass
	Equipment *Equipment
}

func (p Person) String() string {
	return fmt.Sprintf("%s, level: %d", p.Nickname, *p.Level.currentLevel)
}

func CreatePlayer(nickname string) *Person {
	baseClass := GetBaseClass()
	bsLevel := GetStartLevel()
	return &Person{
		Nickname: nickname,
		Level:    bsLevel,
		Class:    baseClass,
	}
}
