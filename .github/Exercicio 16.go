package main

import (
	"fmt"
	"strings"
)

func main() {

	var x, y string

	fmt.Print("Escreva duas palavras: ")
	fmt.Scan(&x, &y)

	if strings.Contains(x, y) {
		fmt.Println("A segunda é uma substring")
	} else {
		fmt.Println("A segunda não é uma substring")
	}
}
