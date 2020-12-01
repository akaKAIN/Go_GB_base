package main

import (
	"os"
	"testing"
)

func TestInput(t *testing.T) {
	var (
		endInput        = []byte("\n")
		emptyInput      = []byte("")
		expressionInput = []byte("2+10")
	)

	//Создаем пайп для передачи данных в терминал и чтения
	r, w, err := os.Pipe()
	defer func() {
		if err := w.Close(); err != nil {
			t.Error("Error closing writer: ", err)
		}
	}()
	if err != nil {
		t.Fatalf("Pipe error: %s", err)
	}

	confirmInput := func() {
		_, err = w.Write(endInput)
		if err != nil {
			t.Fatalf("Write error: %s", err)
		}
	}

	// Передаем в Пайп пустое значение
	_, err = w.Write(emptyInput)
	if err != nil {
		t.Fatalf("Write error: %s", err)
	}
	confirmInput()

	stdin := os.Stdin
	defer func() { os.Stdin = stdin }()
	os.Stdin = r

	input, err := Input()
	if input != "" {
		t.Fatalf("Empty input from user is unacceptable")
	}

	_, err = w.Write(expressionInput)
	if err != nil {
		t.Fatalf("Write expretion error: %s", err)
	}
	confirmInput()

	input, err = Input()
	if err != nil {
		t.Fatalf("No error expected, but got: %s", err)
	}
	if input != string(expressionInput) {
		t.Fatalf("Expected input: %s, but got %s", string(expressionInput), input)
	}
}
