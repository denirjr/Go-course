package manipulador

import "html/template"

//Modelos armazena os modelos de pagina que irá ser execultado pelos manipuladores
var Modelos = template.Must(template.ParseFiles("html/ola.html"))
