package models

const (
	zero = iota
	basePoint
	baseTen    = iota * 10
	adventurer = "Adventurer"
	mage       = "Mage"
	warrior    = "Warrior"
)

type GameClass struct {
	Title      string
	ClassProps Properties
}

func GetBaseClass() *GameClass {
	return &GameClass{
		Title: adventurer,
		ClassProps: Properties{
			Mind:   basePoint,
			Might:  basePoint,
			Speed:  baseTen,
			Health: baseTen,
			Mana:   zero,
		},
	}
}
