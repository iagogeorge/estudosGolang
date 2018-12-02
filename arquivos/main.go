package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {

	arquivo, err := os.Open("cidades.csv")

	if err != nil {
		fmt.Println("[main] Houve um erro ao abrir arquivo. Erro: ", err.Error())
		return
	}

	/*scanner := bufio.NewScanner(arquivo)
	for scanner.Scan() {
		linha := scanner.Text()
		fmt.Println("o conteudo da linha é: ", linha)
	}*/

	leitorCsv := csv.NewReader(arquivo)
	conteudo, err1 := leitorCsv.ReadAll()
	if err != nil {
		fmt.Println("[main] Houve um erro ao abrir arquivo. Erro: ", err1.Error())
		return
	}
	for indice, valor := range conteudo {
		fmt.Printf("indice[%d] = conteudo %s\r\n", indice, valor)
	}

	arquivo.Close()
}
