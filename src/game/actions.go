package game

import (
	"github.com/akaKAIN/Go_GB_base/src/game/models"
	"github.com/akaKAIN/Go_GB_base/src/input"
)

func StartGame() {

}

func GetPlayer() *models.Person {
	name := input.Input("Enter name of your hero: ")
	return models.CreatePlayer(name)
}
