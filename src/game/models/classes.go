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
			Might:  0,
			Speed:  0,
			Health: 0,
			Mana:   baseTen,
		},
	}
}

func GetWarriorClass() *GameClass {
	return &GameClass{
		Title: warrior,
		ClassProps: Properties{
			Mind:   0,
			Might:  basePoint,
			Speed:  0,
			Health: baseTen * 2,
			Mana:   0,
		},
	}
}
