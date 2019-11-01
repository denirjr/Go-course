package main

import (
	"fmt"

	"github.com/denirjr/funcoes_basico/matematica"
)

func main() {
	resultado := matematica.Calculo(matematica.Multiplicador, 2, 2)
	fmt.Printf("o resultado da multiplicacao foi: %d\r\n", resultado)
	resultado = matematica.Soma(3, 3)
	fmt.Printf("o resultado da soma foi: %d\r\n", resultado)
	resultado = matematica.Calculo(matematica.Divisor, 12, 3)
	fmt.Printf("o resultado da divisao foi: %d\r\n", resultado)
	resultado, resto := matematica.DivisorComResto(12, 5)
	fmt.Printf("o resultado da divisao foi %d e o resto foi: %d\r\n", resultado, resto)
}
