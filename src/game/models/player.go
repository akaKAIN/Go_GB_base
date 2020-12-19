package models

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

type Calculator interface {
	Calc()
}

type Person struct {
	Nickname      string
	Level         *Level
	Class         *GameClass
	Equipment     *Equipment
	SelfProps     *Properties
	TotalProps    *Properties
	ContextLogger *logrus.Entry
}

func (p Person) String() string {
	return fmt.Sprintf("%s, level: %d", p.Nickname, *p.Level.currentLevel)
}

func (p *Person) refreshTopProps() {
	p.TotalProps = CalcProps(*p.SelfProps, *p.Equipment.GetEquipmentTotalProps(), p.Class.ClassProps)
}

func CreatePlayer(nickname string) *Person {
	bsClass := GetBaseClass()
	bsLevel := GetStartLevel()
	bsEquipment := CreateEmptyEquipment()
	bsSelfProps := Properties{
		Mind:   basePoint,
		Might:  basePoint,
		Speed:  basePoint * baseTen,
		Health: basePoint * baseTen,
		Mana:   0,
	}
	person := Person{
		Nickname:  nickname,
		Level:     bsLevel,
		Class:     bsClass,
		Equipment: bsEquipment,
		SelfProps: &bsSelfProps,
	}
	person.refreshTopProps()
	return &person
}
