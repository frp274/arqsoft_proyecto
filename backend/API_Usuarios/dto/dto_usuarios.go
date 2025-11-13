package dto

type UsuarioDto struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
	EsAdmin  bool   `json:"es_admin"`
}

type UsuariosDto []UsuarioDto

type LoginResponse struct {
	Id      int    `json:"id"`
	Token   string `json:"token"`
	EsAdmin bool   `json:"es_admin"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreateUsuarioRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Nombre   string `json:"nombre" binding:"required"`
	Apellido string `json:"apellido" binding:"required"`
	Password string `json:"password" binding:"required"`
	EsAdmin  bool   `json:"es_admin"`
}
