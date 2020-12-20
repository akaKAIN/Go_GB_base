package models

import "fmt"

type Properties struct {
	Mind   int16
	Might  int16
	Speed  int16
	Health int16
	Mana   int16
}

func (p Properties) String() string {
	return fmt.Sprintf(
		"\thealth: %v\n\tmana: \t%v\n\tspeed: \t%v\n\tmind: \t%v\n\tmight: \t%v\n",
		p.Health, p.Mana, p.Speed, p.Mind, p.Might,
	)
}

func CalcProps(props ...Properties) *Properties {
	resultProp := Properties{
		Mind:   zero,
		Might:  zero,
		Speed:  zero,
		Health: zero,
		Mana:   zero,
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
