package main

import "fmt"

// Imovel é uma struct que armazena dados do imovel
type Imovel struct {
	x     int
	y     int
	Nome  string
	valor int
}

func main() {

	casa := new(Imovel)
	fmt.Printf("Casa é: %p - %+v\r\n", &casa, casa)

	chacara := Imovel{17, 28, "chacara linda", 280000}
	fmt.Printf("Chacara é: %p - %+v\r\n", &chacara, chacara)
	mudaImovel(&chacara)
	fmt.Printf("Chacara é: %p - %+v\r\n", &chacara, chacara)
}

func mudaImovel(imovel *Imovel) {

	imovel.x = imovel.x + 10
	imovel.y = imovel.y - 5

}
