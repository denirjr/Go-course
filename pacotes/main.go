package main

import (
	"fmt"

	"github.com/denirjr/gocourse/pacotes/operadora"
	"github.com/denirjr/gocourse/pacotes/prefixo"
)

//NomeDoUsuario é o nome do usuario do sistema
var NomeDoUsuario = "Denir"

func main() {

	fmt.Printf("Nome do usuario: %s\r\n", NomeDoUsuario)
	fmt.Printf("Prefixo da capital: %d\r\n", prefixo.Capital)
	fmt.Printf("Nome da operadora: %s\r\n", operadora.NomeOperadora)

}
