package main

import (
	"github.com/akaKAIN/Go_GB_base/src/myreader"
	"log"
	"net/url"
)

func main() {
	path := "conf.yaml"
	config, err := myreader.GetSimpleConfigYAML(path)
	if err != nil {
		log.Printf("Get еconfig error: %s", err)
	}
	log.Println(config)
	res, err := url.Parse(config.DBUrl)
	if err != nil {
		log.Println(err)
	}
	log.Println(res.Port())
	log.Println(res)
}
