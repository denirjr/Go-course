package model

import "errors"

// Imovel representações de informações de um imóvel
type Imovel struct {
	X     int    `json:"coordenada_x"`
	Y     int    `json:"coordenada_y"`
	Nome  string `json:"nome"`
	valor int
}

// ErrValorDeImovelInvalido é um erro apresentado quando é atribuido a um imovél com valor muito baixo
var ErrValorDeImovelInvalido = errors.New("O valor informado não é valido para um imovél")

// ErrValorDeImovelMuitoAlto é um erro apresentado quando é atribuido a um imovel com valor miuto alto
var ErrValorDeImovelMuitoAlto = errors.New("O valor informado é muito alto")

//GetValor retorna o valor do imovel
func (i *Imovel) GetValor() int {
	return i.valor
}

//SetValor defini o valor do imovel
func (i *Imovel) SetValor(valor int) (err error) {
	err = nil
	if valor < 50000 {
		err = ErrValorDeImovelInvalido
		return
	} else if valor > 100000 {
		err = ErrValorDeImovelMuitoAlto
		return
	}
	i.valor = valor
	return
}
