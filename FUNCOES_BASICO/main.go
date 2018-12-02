package main

import (
	"cursoDeGo/FUNCOES_BASICO/matematica"
	"fmt"
)

func main() {

	resultado := matematica.Calculo(matematica.Multiplicador, 8, 2)
	fmt.Printf("o resultado foi: %d\r\n", resultado)

	resultado = matematica.Soma(2, 3)
	fmt.Printf("o resultado da soma foi: %d\r\n", resultado)

	resultado = matematica.Calculo(matematica.Divisor, 12, 3)
	fmt.Printf("o resultado da divisao foi: %d\r\n", resultado)

	resultado, resto := matematica.DivisorComResto(12, 5)
	fmt.Printf("o resultado da divisao foi: %d -> %d ", resultado, resto)

}
