package models

import "fmt"

type Properties struct {
	Mind   int16
	Might  int16
	Speed  int16
	Health int16
	Mana   int16
}

func (p *Properties) String() string {
	return fmt.Sprintf(
		"health: %v\nmana: %v\nspeed: %v\nmind: %v\nmight: %v\n",
		p.Health, p.Mana, p.Speed, p.Mind, p.Might,
	)
}

func CalcProps(props ...Properties) *Properties {
	resultProp := Properties{
		Mind:   0,
		Might:  0,
		Speed:  0,
		Health: 0,
		Mana:   0,
	}

	for _, prop := range props {
		resultProp.Health += prop.Health
		resultProp.Mana += prop.Mana
		resultProp.Might += prop.Might
		resultProp.Mind += prop.Mind
		resultProp.Speed += prop.Speed
	}
	return &resultProp
}
