package main

import "fmt"

var (
	//Endereco é um valor importante e tem que ser público
	Endereco string
	//Telefone é um valor importante e tem que ser público
	Telefone            string
	quantidade, estoque int
	comprou             bool
	valor               float64
	palavras            rune
)

func main() {
	teste := "Valor teste"
	fmt.Printf("endereco %s\r\n", Endereco)
	fmt.Printf("quantidade %d\r\n", quantidade)
	fmt.Printf("comprou %v\r\n", comprou)
	fmt.Printf("palavras %v\r\n", palavras)
	fmt.Printf("o valor de teste é %v\r\n", teste)

}
