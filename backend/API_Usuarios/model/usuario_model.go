package model

type Usuario struct {
	Id           int    `gorm:"primaryKey"`
	Username     string `gorm:"type:varchar(100);not null;unique"`
	Email        string `gorm:"type:varchar(255);not null;unique"`
	Nombre       string `gorm:"type:varchar(100);not null"`
	Apellido     string `gorm:"type:varchar(100);not null"`
	PasswordHash string `gorm:"type:varchar(64);not null"`
	EsAdmin      bool   `gorm:"column:es_admin;not null;default:false"`
}

type Usuarios []Usuario
//CONFIRMAR ACTIVIDAD / AVISAR X MAIL CUANDO SE CARGA BIEN UN USUARIO
//AVISAR SUSCRIPCION EXITOSA A ACTIVIDAD POR MAIL
//