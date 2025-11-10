package model

type Usuario struct {
	Id              int    `gorm:"primaryKey"`
	Nombre_apellido string `gorm:"type:varchar(255);not null"`
	UserName        string `gorm:"column:Username;type:varchar(100);not null;unique"`
	Es_admin        bool   `gorm:"not null;default:false"`
	PasswordHash    string `gorm:"type:varchar(64);not null"`
}

type Usuarios []Usuario
