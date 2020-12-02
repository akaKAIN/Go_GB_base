package main

import (
	"os"
	"testing"
)

func TestInput(t *testing.T) {
	var (
		titleMsg        = "Test input"
		endInput        = []byte("\n")
		expressionInput = []byte("2df0")
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

	// Передаем в Пайп значение
	_, err = w.Write(expressionInput)
	if err != nil {
		t.Fatalf("Write error: %s", err)
	}
	confirmInput()

	stdin := os.Stdin
	defer func() { os.Stdin = stdin }()
	os.Stdin = r

	input := Input(titleMsg)
	if input != string(expressionInput) {
		t.Fatalf("Expected input %s, but got %s", string(expressionInput), input)
	}
}

// TODO
func Test_makeExpression(t *testing.T) {}

// TODO
func Test_getOperand(t *testing.T) {}

// TODO
func Test_getHandlerByOperator(t *testing.T) {}
