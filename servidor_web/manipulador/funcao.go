package manipulador

import (
	"fmt"
	"net/http"
)

// Funcao é um manipulador de função http na rota /funcao
func Funcao(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Aqui é um manipulador de funcao num pacote")
}
