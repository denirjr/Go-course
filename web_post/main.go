package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/denirjr/gocourse/web_post/model"
)

func main() {

	// Cria client
	cliente := &http.Client{
		Timeout: time.Second * 30,
	}

	// Define dados a se enviar
	usuario := model.Usuario{}
	usuario.ID = 1
	usuario.Nome = "Denir Damasceno Jr."

	// ultiliza o Marshal do pacote json para transformar os dados em json
	coteudoEnviar, err := json.Marshal(usuario)
	if err != nil {
		fmt.Println("[main] Erro ao gerar o json do objeto usuario. Erro: ",
			err.Error())
		return
	}

	// Cria um novo request do tipo post com endpoint desejado e cria um reader (bytes.NewBuffer) p com o conteudo a se enviar
	request, err := http.NewRequest("POST", "http://requestbin.net/r/yxx5ewyx", bytes.NewBuffer(coteudoEnviar))
	if err != nil {
		fmt.Println("[main] Erro ao criar um request para o resquestbin. Erro: ",
			err.Error())
		return
	}

	// define os itens da autenticação a se enviar e execulta a ação (cliente.Do)
	request.SetBasicAuth("fizz", "buzz")
	request.Header.Set("content-type", "application/json; charset=utf-8")
	resposta, err := cliente.Do(request)

	if err != nil {
		fmt.Println("[main] Erro ao chamar o serviço do request bin . Erro: ",
			err.Error())
		return
	}
	defer resposta.Body.Close()

	// Se Trata o conteúdo
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[main] Erro ao ler o conteudo retornado pelo request bin: ",
				err.Error())
			return
		}
		fmt.Println(" ")
		fmt.Println(string(corpo))
	}
}
