package operadora

import (
	"strconv"

	"github.com/denirjr/gocourse/pacotes/prefixo"
)

//NomeOperadora representa o nome da operadora
var NomeOperadora = "Telecom"

//PrefixoDaCapitalOPeradora mais o nome da operadora
var PrefixoDaCapitalOPeradora = strconv.Itoa(prefixo.Capital) + " " + NomeOperadora
