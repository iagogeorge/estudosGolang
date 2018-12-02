package main

import (
	"fmt"
)

func main() {

	var nun []int

	fmt.Println(nun, len(nun), cap(nun))

	nun = make([]int, 5)

	fmt.Println(nun, len(nun), cap(nun))

	capitais := []string{"lisboa"}

	fmt.Println(capitais, len(capitais), cap(capitais))

	capitais = append(capitais, "Recife")

	fmt.Println(capitais, len(capitais), cap(capitais))

	cidades := make([]string, 4)

	cidades[0] = "Recife"
	cidades[1] = "Joao Pessoa"
	cidades[2] = "Natal"
	cidades[3] = "Fortaleza"

	fmt.Println(cidades, len(cidades), cap(cidades))

	capitais[1] = "recife"

	fmt.Println(capitais, len(capitais), cap(capitais))

	for index, valor := range cidades {

		fmt.Printf("indice[%d] = Valor[%s]\r\n", index, valor)

	}

	capitaisBrasil := cidades[2:4]

	fmt.Println(capitaisBrasil, len(capitaisBrasil), cap(capitaisBrasil))

}
