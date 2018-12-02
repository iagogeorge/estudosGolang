package main

import (
	"fmt"
)

func main() {

	meses := 0
	situacao := true
	cidade := "Recife"

	if meses <= 6 {
		fmt.Println("esse credor deve há pouco tempo")
	}

	if situacao {
		fmt.Println("ok")
	}

	if cidade != "Olinda" {
		fmt.Println("outra cidade")

	}

	if descricao, status := haQuantoTempoEstaDevendo(meses); status {
		fmt.Println(descricao)
		return
	}

	fmt.Println("obrigado!")
}

func haQuantoTempoEstaDevendo(meses int) (descricao string, status bool) {

	if meses > 0 {
		status = true
		descricao = "o cliente esta devendo"
		return
	}

	descricao = "o cliente ta de boa"

	return
}
