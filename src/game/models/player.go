package models

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

type Calculator interface {
	Calc()
}

type Person struct {
	Nickname        string
	Level           *Level
	Class           *GameClass
	Equipment       *Equipment
	SelfProps       *Properties
	TotalProps      *Properties
	ContextLogger   *logrus.Entry
	HealthIndicator *LifeIndicator
	ManaIndicator   *LifeIndicator
}

func (p Person) String() string {
	return fmt.Sprintf("%s, level: %d", p.Nickname, *p.Level.currentLevel)
}

func (p Person) ShowInfo() string {
	var message string
	message = fmt.Sprintf("player: %s,\t%v lvl\t%s\t%s\nprops:\n%s",
		p.Nickname, p.Level.GetCurrentLevel(), p.HealthIndicator, p.ManaIndicator, p.TotalProps)
	return message
}

func (p *Person) RefreshTopProps() {
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
		Mana:   zero,
	}
	person := Person{
		Nickname:  nickname,
		Level:     bsLevel,
		Class:     bsClass,
		Equipment: bsEquipment,
		SelfProps: &bsSelfProps,
	}

	person.RefreshTopProps()
	person.InitLogger()
	person.HealthIndicator = GetNewLifeIndicator("Health", uint16(person.TotalProps.Health), true)
	person.ManaIndicator = GetNewLifeIndicator("Mana", uint16(person.TotalProps.Mana), false)
	person.LogAction("I'm was born!")
	return &person
}

func (p *Person) InitLogger() {
	p.ContextLogger = logrus.WithFields(
		logrus.Fields{
			"name": p.Nickname,
			"lvl":  fmt.Sprintf("%v", p.Level.GetCurrentLevel()),
		},
	)
}

func (p *Person) LogAction(actionText string) {
	p.ContextLogger.Info(actionText)
}

func (p *Person) Attack(enemy *Person) {
//	TODO Attack method here or in Item struct (need to fink)
}
