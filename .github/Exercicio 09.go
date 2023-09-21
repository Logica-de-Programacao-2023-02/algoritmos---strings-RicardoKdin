package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	var x string
	var y string
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Escreva uma frase: ")
	scanner.Scan()
	frase := scanner.Text()

	fmt.Println("Escreva uma letra que quer adicionar: ")
	fmt.Scan(&x)

	fmt.Println("Escreva a letra que quer retirar: ")
	fmt.Scan(&y)

	z := strings.ReplaceAll(frase, y, x)

	fmt.Println(z)

}
