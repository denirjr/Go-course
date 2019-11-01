package matematica

//Calculo que execulta qualquer tipo de calculo, basta enviar a funcao desejada
func Calculo(funcao func(int, int) int, numeroA int, numeroB int) int {
	return funcao(numeroA, numeroB)
}

//Multiplicador miltiplica X vezes o y
func Multiplicador(x int, y int) int {
	return x * y
}

//Divisor efetual a divisao do numeroA pleo numeroB
func Divisor(numeroA int, numeroB int) (resultado int) {
	resultado = numeroA / numeroB
	return
}

// DivisorComResto retorna o resultado e o resto da divisão
func DivisorComResto(numeroA int, numeroB int) (resultado int, resto int) {
	resultado = numeroA / numeroB
	resto = numeroA % numeroB
	return
}
