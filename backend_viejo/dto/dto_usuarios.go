package dto

type UsuarioDto struct {
	Id       int    `json:"id"`
	UserName string `json:"Username"`
}

type UsuariosDto []UsuarioDto

type LoginResponse struct {
	Id    int    	`json:"id"`
	Token string 	`json:"token"`
	Es_admin bool	`json:"es_admin"`
}

type LoginRequest struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
}
