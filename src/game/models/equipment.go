package models

type Equipment struct {
	RightHand *Hand
	Body      *Body
}

type Hand struct {
	Item *Item
}

type Body struct {
	Item *Item
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
