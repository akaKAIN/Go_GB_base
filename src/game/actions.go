package game

import (
	"github.com/akaKAIN/Go_GB_base/src/game/models"
	"github.com/akaKAIN/Go_GB_base/src/input"
)


func getPlayer() *models.Person {
	name := input.Input("Enter name of your hero: ")
	return models.CreatePlayer(name)
}


func StartGame() {
	getPlayer()
}
