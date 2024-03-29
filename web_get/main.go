package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	cliente := &http.Client{
		Timeout: time.Second * 30,
	}

	resposta, err := cliente.Get("http://www.google.com.br")
	if err != nil {
		fmt.Println("[main] Erro ao pesquisar a pagina do google Brasil. Erro: ",
			err.Error())
		return
	}
	defer resposta.Body.Close()

	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[main] Erro ao ler o conteudo  da pagina do google Brasil. Erro: ",
				err.Error())
			return
		}
		fmt.Println(string(corpo))
	}

	request, err := http.NewRequest("GET", "http://www.google.com.br", nil)
	if err != nil {
		fmt.Println("[main] Erro ao criar um request para o google brasil. Erro: ",
			err.Error())
		return
	}
	request.SetBasicAuth("teste", "teste")
	resposta, err = cliente.Do(request)

	//Should create function
	if err != nil {
		fmt.Println("[main] Erro ao pesquisar a pagina do google Brasil. Erro: ",
			err.Error())
		return
	}
	defer resposta.Body.Close()

	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[main] Erro ao ler o conteudo  da pagina do google Brasil. Erro: ",
				err.Error())
			return
		}
		fmt.Println(" ")
		fmt.Println(string(corpo))
	}
}
