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

func createEnemy(name string) *models.Person {
	return models.CreatePlayer(name)
}

func StartGame() {
	player := getPlayer()
	sword := models.CreateWeapon("Base", "none")

	models.EquipItem(player.Equipment.Hand, sword)
	player.RefreshTopProps()
	player.LogAction(fmt.Sprintf("%s: %+v was equipped", sword.Title, sword.Properties))
	fmt.Println(player.ShowInfo())

	worm := createEnemy("Worm")
	player.PhysicalAttack(worm)

	for worm.HealthIndicator.Get() > 0 {
		player.PhysicalAttack(worm)
	}
}
