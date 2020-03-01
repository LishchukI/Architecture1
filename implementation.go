package lab1

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/LishchukI/Architecture1/stack"
)

func isOperator(str string) bool {
	switch str {
	case "+":
		return true
	case "-":
		return true
	case "/":
		return true
	case "*":
		return true
	case "^":
		return true
	default:
		return false
	}
}

func isValidNumber(str string) bool {
	_, e := strconv.Atoi(str)
	if e != nil || strings.Contains(str, "+") {
		return false
	}
	return true
}

// PrefixToInfix comment
func PrefixToInfix(input string) (string, error) {
	prefixExp := strings.Fields(input)
	s := stack.NewStack()
	length := len(prefixExp)

	for i := length - 1; i >= 0; i-- {
		symbol := prefixExp[i]

		switch {

		case isOperator(symbol):
			operand1Raw, err := s.Pop()
			operand2Raw, err := s.Pop()

			if err != nil {
				return "", fmt.Errorf("Invalid prefix expression")
			}

			operand1, _ := operand1Raw.(string)
			operand2, _ := operand2Raw.(string)

			temp := "(" + strings.Join([]string{operand1, symbol, operand2}, " ") + ")"

			s.Push(temp)

		case isValidNumber(symbol):
			s.Push(symbol)

		default:
			return "", fmt.Errorf("Invalid symbol in prefix expression")
		}
	}

	infixExpRaw, err := s.Pop()

	if err != nil || s.Size() != 0 {
		return "", fmt.Errorf("Invalid prefix expression")
	}

	infixExp, _ := infixExpRaw.(string)

	return infixExp, nil
}
