package structs

import (
	"fmt"
)

//Imovel teste
type Imovel struct {
	X     int
	Y     int
	Nome  string
	Valor int
}

func main() {

	casa := Imovel{}
	fmt.Printf("%#v\r", casa)

	apartamento := Imovel{17, 56, "meu AP", 780000}
	fmt.Printf("%#v\r", apartamento)

	mudaImovel(&apartamento)
	//apartamento := Imovel{17, 56, "meu AP", 780000}
	fmt.Printf("%#v\r", apartamento)

}

func mudaImovel(imovel *Imovel) {

	imovel.X = imovel.X + 10
	imovel.Y = imovel.Y - 1
}
