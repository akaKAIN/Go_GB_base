package game

import (
	"fmt"
	"github.com/akaKAIN/Go_GB_base/src/game/models"
	"github.com/akaKAIN/Go_GB_base/src/input"
)

func GetPlayer() {
	name := input.Input("Enter name of your hero: ")
	player := models.CreatePlayer(name)
	fmt.Println(player)

}
