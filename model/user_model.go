package model

type Usuario struct {
	ID       int    `gorm:"primaryKey"` //`json:""` sirve para que al correr el programa no aparezca en mayusculas, y si direct queremos ponerlo con mayus pasa a ser un atributo privado
	Nombre     string `gorm:"not null"`
	Es_admin bool   `gorm:"not null"`
	Password string `gorm:"not null"`
}

type Usuarios []Usuario
