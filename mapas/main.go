package main

import (
	structs "cursoDeGo/FUNCOES_BASICO/structs_basico"
	"fmt"
)

var cache map[string]structs.Imovel

func main() {

	cache = make(map[string]structs.Imovel, 0)

	apto := structs.Imovel{}

	apto.Nome = "apto2"
	apto.X = 20
	apto.Y = 30
	apto.Valor = 70000

	cache["apto1"] = apto

	imovel, achou := cache["apto1"]

	fmt.Println(len(cache))

	if achou {

		fmt.Printf("%v\r\n", imovel)

	}

}
