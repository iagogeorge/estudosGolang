package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {

		fmt.Println(i)

	}

	valor := 0
	teste := true

	for teste {
		valor++

		if valor%5 == 0 {
			teste = false
		}

		fmt.Println("o numero é:", valor)
	}

	for {
		valor--
		fmt.Println("o numero é:", valor)

		if valor == 0 {
			break
		}
	}

	texto := "eu amo go"

	for i, v := range texto {

		fmt.Printf("%d\r\n - > %s\r\n", i, string(v))

	}

}
