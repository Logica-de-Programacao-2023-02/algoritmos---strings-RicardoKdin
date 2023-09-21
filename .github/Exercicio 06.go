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

	x := strings.Fields(frase)
	y := len(x)

	fmt.Printf("A frase possui %d palavras.", y)

}
