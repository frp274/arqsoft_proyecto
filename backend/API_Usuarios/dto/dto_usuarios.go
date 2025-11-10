package dto

type UsuarioDto struct {
	Id              int    `json:"id"`
	Nombre_apellido string `json:"nombre_apellido"`
	UserName        string `json:"username"`
	Es_admin        bool   `json:"es_admin"`
}

type UsuariosDto []UsuarioDto

type LoginResponse struct {
	Id       int    `json:"id"`
	Token    string `json:"token"`
	Es_admin bool   `json:"es_admin"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreateUsuarioRequest struct {
	Nombre_apellido string `json:"nombre_apellido" binding:"required"`
	UserName        string `json:"username" binding:"required"`
	Password        string `json:"password" binding:"required"`
	Es_admin        bool   `json:"es_admin"`
}
