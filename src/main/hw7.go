package main

import (
	"fmt"
	"github.com/akaKAIN/Go_GB_base/src/myreader"
	"log"
)

func main() {
	path := "conf.yaml"
	config, err := myreader.GetSimpleConfigYAML(path)
	if err != nil {
		log.Printf("Get еconfig error: %s", err)
	}
	validRes := config.GetValidationResult()
	fmt.Println(validRes)
}