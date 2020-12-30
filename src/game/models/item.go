package models

import "fmt"

type Item struct {
	Title      string
	Properties Properties
}

func (i Item) String() string {
	return fmt.Sprintf("%q with props: %+v", i.Title, i.Properties)
}

func CreateWeapon(title string, speciality string) *Item {
	var (
		weaponProps = new(Properties)
	)
	switch speciality {
	case mage:
		title += " Sword"
		weaponProps.Mind = basePoint * 2
		weaponProps.Mana = baseTen * 2
	case warrior:
		title += " staff"
		weaponProps.Might = basePoint * 2
		weaponProps.Speed = basePoint * 2
		weaponProps.Health = baseTen
	default:
		title += " branch"
		weaponProps.Might = basePoint
	}
	return &Item{
		Title:      title,
		Properties: *weaponProps,
	}
}