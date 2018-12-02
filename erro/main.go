package main

import (
	"cursoDeGo/FUNCOES_BASICO/structs_basico"
	"encoding/json"
	"fmt"
)

func main() {

	casa := structs.Imovel{}

	casa.Nome = "casa de iago"
	casa.X = 10
	casa.Y = 20
	casa.Valor = 100000

	fmt.Printf("Casa é %+v\r\n", casa)

	objJSON, _ := json.Marshal(casa)

	fmt.Println("Casa JSON: ", string(objJSON))

}
