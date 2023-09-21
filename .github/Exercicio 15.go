package main

import (
	"fmt"
	"strings"
)

func main() {

	var x string

	fmt.Print("Escreva uma palavra: ")
	fmt.Scan(&x)

	vogais := []string{"A", "a", "E", "e", "I", "i", "O", "o", "U", "u"}

	for _, vogal := range vogais {
		x = strings.ReplaceAll(x, vogal, "*")
	}
	fmt.Println(x)
}
