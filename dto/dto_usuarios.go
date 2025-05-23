package dto

type UsuarioDto struct {
	Id       int    `json:"id"`
	Mail string `json:"mail"`
}

type UsuariosDto []UsuarioDto