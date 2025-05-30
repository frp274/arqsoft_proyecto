package model

type Usuario struct {
	Id       int    	`gorm:"primaryKey"` //`json:""` sirve para que al correr el programa no aparezca en mayusculas, y si direct queremos ponerlo con mayus pasa a ser un atributo privado
	Nombre_apellido     string 	`gorm:"not null"`
	UserName string 		`gorm:"not null"`
	Es_admin bool   	`gorm:"not null"`
	PasswordHash string 	`gorm:"not null"`
}

type Usuarios []Usuario
