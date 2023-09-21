package main

import (
	"fmt"
	"strconv"
)

func main() {

	var x string

	fmt.Print("Digite um número: ")
	fmt.Scan(&x)

	i, err := strconv.ParseFloat(x, 64)

	if err == nil {
		fmt.Println("O valor é: ", i)
	} else {
		fmt.Println("O número não é válido")
	}
}
