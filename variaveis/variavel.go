package main

import (
	"fmt"
)

//Endereco variavel global
var Endereco string
var telefone = "9.9644-1211"
var quantidade int
var comprou bool
var valor float64
var palavras rune

func main() {

	fmt.Printf("endereço: %s\r\n", Endereco)
	fmt.Printf("quantidade: %d\r\n", quantidade)
	fmt.Printf("comprou: %v\r\n", comprou)
	fmt.Printf("palavras: %v\r\n", palavras)

}
