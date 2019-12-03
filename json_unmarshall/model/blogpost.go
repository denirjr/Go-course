package model

//BlogPost armazena dados de um post do Blog
type BlogPost struct {
	UsuarioID int    `json:"userId"`
	ID        int    `json:"ID"`
	Titulo    string `json:"title"`
	Texto     string `json:"body"`
}
