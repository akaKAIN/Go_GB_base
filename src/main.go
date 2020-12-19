package main

import (
	"fmt"
	"github.com/akaKAIN/Go_GB_base/src/game"
	log "github.com/sirupsen/logrus"
)

func main() {
	p := game.GetPlayer()
	contextLogger := log.WithFields(log.Fields{
		"name": p.Nickname,
		"lvl": fmt.Sprintf("%v", p.Level.GetCurrentLevel()),
		"props": fmt.Sprintf("%+v", p.TotalProps),
	})

	contextLogger.Info("I'll be logged with common and other field")
	contextLogger.Info("Me too")
}
