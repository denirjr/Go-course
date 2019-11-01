package main

import (
	"encoding/json"
	"fmt"

	"github.com/denirjr/gocourse/imovel.go-erro/model"
)

func main() {

	casa := model.Imovel{}
	casa.Nome = "Casa Amarela"
	casa.X = 18
	casa.Y = 25
	if err := casa.SetValor(11000000); err != nil {
		fmt.Println("[main] Houve um erro na atribuicão de um valor a casa: ", err.Error())
		if err == model.ErrValorDeImovelMuitoAlto {
			fmt.Println("Corretor verifique a sua avaliação")
		}
		return
	}

	fmt.Printf("Casa é: %+v\r\n", casa)

	objJSON, err := json.Marshal(casa)
	if err != nil {
		fmt.Println("[main] houve um erro na geração do objeto json", err.Error())
		return
	}
	fmt.Println("A casa em json: ", string(objJSON))
}
