package main

import (
	"fmt"
	"regexp"
)

func main() {

	var x string

	fmt.Print("Digite uma string: ")
	fmt.Scan(&x)

	if numeros(x) {
		fmt.Println("A string contém apenas números.")
	} else {
		fmt.Println("A string não contém apenas números.")
	}
}

func numeros(s string) bool {
	match, _ := regexp.MatchString("^[0-9]+$", s)
	return match
}
