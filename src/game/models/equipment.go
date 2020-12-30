package models

type Equipmenter interface {
	Equip(*Item)
}

type Equipment struct {
	Hand *HandWeapon
	Body *BodyArmor
}

func EquipItem(e Equipmenter, item *Item) {
	e.Equip(item)
}

type HandWeapon struct {
	Item *Item
}

type BodyArmor struct {
	Item *Item
}

func (h *HandWeapon) Equip(item *Item) {
	h.Item = item
}

func (b *BodyArmor) Equip(item *Item) {
	b.Item = item
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
