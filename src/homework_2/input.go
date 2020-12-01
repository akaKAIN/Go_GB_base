package main

import (
	"bufio"
	"fmt"
	"os"
)

func Input() (string, error) {
	var scanner = bufio.NewScanner(os.Stdin)
	fmt.Print("Введите простое математическое выражение: ")
	scanner.Scan()
	text := scanner.Text()
	if text == "" {
		return text, fmt.Errorf("Вы ввели пустую строку\n")
	}

	return text, nil
}
