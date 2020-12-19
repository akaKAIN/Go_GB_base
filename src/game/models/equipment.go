package models

import "fmt"

type Equipmenter interface {
	Equip(*Item)
}

type Equipment struct {
	Hand *HandWeapon
	Body *BodyArmor
}

func (e *Equipment) EquipBody(item *Item) {
	//TODO: добавить интерфейс для добавления экипировки
	e.Body.Equip(item)
}

type HandWeapon struct {
	Item *Item
}

func (h *HandWeapon) Equip(item *Item) {
	h.Item = item
}

type BodyArmor struct {
	Item *Item
}

func (b *BodyArmor) Equip(item *Item) {
	b.Item = item
}

type Item struct {
	Title      string
	Properties Properties
}

func (i Item) String() string {
	return fmt.Sprintf("%q with props: %+v", i.Title, i.Properties)
}

func CreateEmptyEquipment() *Equipment {
	return &Equipment{
		Hand: &HandWeapon{Item: new(Item)},
		Body: &BodyArmor{Item: new(Item)},
	}
}

func (e *Equipment) GetEquipmentTotalProps() *Properties {
	var total = new(Properties)
	if e.Body != nil {
		total = CalcProps(*total, e.Body.Item.Properties)
	}
	if e.Hand != nil {
		total = CalcProps(*total, e.Hand.Item.Properties)
	}
	return total
}

func GetWeapon(title string, speciality string) *Item {
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
