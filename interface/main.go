package main

import (
	"fmt"
	"interface/model"
)

func main() {

	jojo := model.Ave{}
	jojo.Nome = "Jojo da silva"

	queroAcordaComUmaGalinha(jojo)
	queroOuvirUmaPataNoLago(jojo)

}

func queroAcordaComUmaGalinha(g model.Galinha) {
	fmt.Println(g.Carcareja())
}

func queroOuvirUmaPataNoLago(g model.Pato) {
	fmt.Println(g.Quack())
}
