package game

import (
	"fmt"
	"github.com/akaKAIN/Go_GB_base/src/game/models"
	"github.com/akaKAIN/Go_GB_base/src/input"
)

func getPlayer() *models.Person {
	name := input.Input("Enter name of your hero: ")
	return models.CreatePlayer(name)
}

func StartGame() {
	player := getPlayer()
	sword := models.GetWeapon("Base", "none")
	fmt.Println(sword)
	player.Equipment.EquipBody(sword)
	player.RefreshTopProps()
	player.LogAction("Sword was equipped")
	fmt.Println(player.ShowInfo())
}
