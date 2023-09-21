package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Escreva uma frase: ")
	scanner.Scan()
	frase := scanner.Text()

	var x string
	var y string

	fmt.Println("Insira um caractere que quer adicionar: ")
	fmt.Scan(&x)
	fmt.Println("Insira um caractere que quer retirar: ")
	fmt.Scan(&y)

	z := strings.ReplaceAll(frase, y, x)

	fmt.Println(z)

}
