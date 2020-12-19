package models

const (
	zero = iota
	basePoint
	baseTen    = 10
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
			Mind:   zero,
			Might:  zero,
			Speed:  zero,
			Health: zero,
			Mana:   zero,
		},
	}
}
