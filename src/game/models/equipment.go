package models

type Equipmenter interface {
	Equip(*Item)
}

type Equipment struct {
	RightHand *HandWeapon
	Body      *BodyArmor
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

func CreateEmptyEquipment() *Equipment {
	return &Equipment{
		RightHand: nil,
		Body:      nil,
	}
}

func (e *Equipment) GetEquipmentTotalProps() *Properties {
	var total = new(Properties)
	if e.Body != nil {
		total = CalcProps(*total, e.Body.Item.Properties)
	}
	if e.RightHand != nil {
		total = CalcProps(*total, e.RightHand.Item.Properties)
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
