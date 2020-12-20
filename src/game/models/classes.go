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

func GetMagClass() *GameClass {
	return &GameClass{
		Title: mage,
		ClassProps: Properties{
			Mind:   basePoint,
			Might:  zero,
			Speed:  zero,
			Health: zero,
			Mana:   baseTen,
		},
	}
}

func GetWarriorClass() *GameClass {
	return &GameClass{
		Title: warrior,
		ClassProps: Properties{
			Mind:   zero,
			Might:  basePoint,
			Speed:  zero,
			Health: baseTen * 2,
			Mana:   zero,
		},
	}
}
