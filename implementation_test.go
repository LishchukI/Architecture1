package lab1

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCasesValidPrefixToInfix(t *testing.T) {
	assert := assert.New(t)

	expressions := []string{" 4 33",
		"+ / - - * 13 8 / 9 1 ^ 6 3 3 ^ - 3 * 7 8 3",
		"* + 11 2 7",
		"100"}

	outcome := []string{"(4 + 33)",
		"(((((13 * 8) - (9 / 1)) - (6 ^ 3)) / 3) + ((3 - (7 * 8)) ^ 3))",
		"((11 + 2) * 7)",
		"100"}

	for i := 0; i < len(expressions); i++ {
		res, err := PrefixToInfix(expressions[i])
		if assert.Nil(err) {
			assert.Equal(res, outcome[i])
		}
	}

}

func TestCasesInvalidPrefixToInfix(t *testing.T) {
	assert := assert.New(t)

	invalid_expressions := []string{"+ 4 2 33",
		"- - * 13 8 / 9 1 ^ 6 3 3 ^ - 3 * 7 8 3",
		"* + 11 + 2 7",
		"* 100"}

	for i := 0; i < len(invalid_expressions); i++ {
		_, err := PrefixToInfix(invalid_expressions[i])
		if assert.NotNil(err) {
			assert.Equal(err.Error(), "Invalid prefix expression")
		}
	}

	invalid_symbols_expressions := []string{"+ a 2 33",
		"- * 13% 8 / 9 1 ^ 6 d 3 ^ - 3 * o 8 3",
		"& + 11 + 2 7",
		"* +100"}

	for i := 0; i < len(invalid_symbols_expressions); i++ {
		_, err := PrefixToInfix(invalid_symbols_expressions[i])
		if assert.NotNil(err) {
			assert.Equal(err.Error(), "Invalid symbol in prefix expression")
		}
	}

}

func ExamplePrefixToInfix() {

	res, err := PrefixToInfix("* + 2 22 - 18 0")
	if err == nil {
		fmt.Println(res)
	}

	// Output:
	// ((2 + 22) * (18 - 0))
}
