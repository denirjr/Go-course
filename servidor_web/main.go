package main

import (
	"fmt"
	"net/http"

	"github.com/denirjr/gocourse/servidor_web/manipulador"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Ola mundo!")
	})
	http.HandleFunc("/funcao", manipulador.Funcao)
	http.HandleFunc("/ola", manipulador.Ola)
	fmt.Println("Estou escutando, comandante ... ")
	http.ListenAndServe(":8181", nil)

}
