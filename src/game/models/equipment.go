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
