package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

func Input(msg string) (string, error) {
	var scanner = bufio.NewScanner(os.Stdin)
	fmt.Print(msg)
	scanner.Scan()
	text := scanner.Text()
	if text == "" {
		return text, fmt.Errorf("Вы ввели пустую строку\n")
	}

	return text, nil
}

func isCleanedString(word string) bool {
	isClean, err := regexp.MatchString("^[0-9]", word)
	if err != nil {
		log.Printf("Cleaned error: %s\n", err)
	}
	return isClean
}
