package matematica

/*
Calculo executa qualquer tipo de calculo basico basta enviar a funçao de desejada
*/
func Calculo(funcao func(int, int) int, numeroA int, numeroB int) int {

	return funcao(numeroA, numeroB)
}

// Multiplicador multiplica
func Multiplicador(x int, y int) int {

	return x * y

}

//Divisor
func Divisor(numeroA int, numeroB int) (resultadoDivisao int) {
	resultadoDivisao = numeroA / numeroB
	return

}

//DivisorComResto
func DivisorComResto(numeroA int, numeroB int) (resultadoDivisao int, resto int) {
	resultadoDivisao = numeroA / numeroB
	resto = numeroA % numeroB
	return

}
