package input

import (
	"bytes"
	"github.com/akaKAIN/Go_GB_base/src/mathoperation"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInput(t *testing.T) {
	userInput := []byte("2df0\n")
	expect := "2df0"

	stdin := new(bytes.Buffer)
	stdin.Write(userInput)
	result := Input(string(userInput), stdin)
	assert.Equal(t, result, expect, "Should be equal")
}

// TODO
func Test_makeExpression(t *testing.T) {}

func Test_getOperand(t *testing.T) {
	tests := []struct {
		in  string
		out int
	}{
		{in: "123\n", out: 123},
		{in: "321\n", out: 321},
	}
	stdin := new(bytes.Buffer)
	for _, tc := range tests {

		stdin.Write([]byte(tc.in))
		num := getOperand(stdin)
		assert.Equal(t, tc.out, num, "Should be equal")

	}
}

func Test_getHandlerByOperator(t *testing.T) {
	m := mathoperation.GetOperationsMap()
	operators := []string{"+", "-", "*", "/"}
	oper1, oper2 := 10, 17
	stdin := new(bytes.Buffer)
	for _, operator := range operators {
		stdin.Write([]byte(operator + "\n"))
		handler := getHandlerByOperator(stdin)
		handlerResult := handler(oper1, oper2)
		expectedResult := m[operator](oper1, oper2)
		assert.Equal(t, handlerResult, expectedResult, "Should be equal")
	}
}
