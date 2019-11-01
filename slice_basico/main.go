package main

import "fmt"

func main() {
	capitais := []string{"Lisboa"}
	// fmt.Println(capitais, len(capitais), cap(capitais))
	capitais = append(capitais, "Brasilia")
	// fmt.Println(capitais, len(capitais), cap(capitais))
	cidades := make([]string, 5)
	cidades[0] = "New york"
	cidades[1] = "London"
	cidades[2] = "Madeira"
	cidades[3] = "Tokio"
	cidades[4] = "Singapur"
	// fmt.Println(cidades, len(cidades), cap(cidades))
	capitais[1] = "Brasília 2"
	fmt.Println(capitais, len(capitais), cap(capitais))
	for indice, cidade := range cidades {
		fmt.Printf("Cidade[%d] = %s\n\r", indice, cidade)
	}
	//Primeiro item  começa com indice 0
	// segundo item começa com indice 1
	capitaisAsia := cidades[3:5]
	fmt.Println(capitaisAsia)
	temp1 := cidades[:2]
	fmt.Println(temp1)

	// Peguei o tamanho do array subtrai por 2 resultou em 3 que é o terceiro indice
	//  passe por fim o ultimo indice como ":"
	temp2 := cidades[len(cidades)-2:]
	fmt.Println(temp2)
	indiceDoItemARetirar := 2
	temp := cidades[:indiceDoItemARetirar]
	temp = append(temp, cidades[indiceDoItemARetirar+1:]...)
	copy(cidades, temp)
	fmt.Println(cidades)
}
