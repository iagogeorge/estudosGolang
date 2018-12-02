package main

import (
	"cursoDeGo/variaveis/pacotes/operadora"
	"cursoDeGo/variaveis/pacotes/prefixo"
	"fmt"
)

//Nome nome do usuario
var Nome = "Iago"

func main() {

	fmt.Printf("nome : %s\r\n", Nome)
	fmt.Printf("prefixo : %d\r\n", prefixo.Capital)
	fmt.Printf("teste: %s\r\n", operadora.PrefixoDaCapitalOperadora)
	fmt.Printf("Nome da Operadora: %s\r\n", operadora.NomeOperadora)
}
