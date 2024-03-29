package manipulador

import (
	"fmt"
	"net/http"
	"time"

	"github.com/denirjr/gocourse/servidor_web/model"
)

// Ola é o manipulador da request  a rota/ola
func Ola(w http.ResponseWriter, r *http.Request) {
	hora := time.Now().Format("15:04:05")
	pagina := model.Pagina{}
	pagina.Hora = hora
	if err := Modelos.ExecuteTemplate(w, "ola.html", pagina); err != nil {
		http.Error(w, "Houve um erro na renderização da pagina.", http.StatusInternalServerError)
		fmt.Println("[Ola] Erro na execulsão desse modelo: ", err.Error())
	}
}
