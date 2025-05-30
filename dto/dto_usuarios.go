package dto

type UsuarioDto struct {
	Id       int    `json:"id"`
	UserName string `json:"username"`
}

type UsuariosDto []UsuarioDto

type LoginResponse struct {
	Id    int    `json:"id"`
	Token string `json:"token"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
