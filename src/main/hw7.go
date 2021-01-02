package main

import (
	"github.com/akaKAIN/Go_GB_base/src/myreader"
	"log"
	"os"
)

func main() {
	path := "conf.yaml"
	config, err := myreader.GetSimpleConfigYAML(path)
	if err != nil {
		log.Printf("Get еconfig error: %s", err)
		os.Exit(1)
	}
	vr := config.GetValidationResult()
	vr.Print()
}
