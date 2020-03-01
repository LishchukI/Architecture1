package main

import (
	"fmt"
	"os"
	"strings"

	lab1 "github.com/LishchukI/Architecture1"
)

func main() {
	output, error := lab1.PrefixToInfix(strings.Join(os.Args[1:], " "))
	if error != nil {
		fmt.Println(error)
	}
	fmt.Println(output)
	fmt.Println("Build version: " + BuildVersion)
}
