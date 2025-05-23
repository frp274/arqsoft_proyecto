package model

type Actividad struct {
	ID          int       `gorm:"primaryKey"`
	Nombre        string    `gorm:"not null"`
	Capacidad    int       `gorm:"not null"`
	Descripcion string    `gorm:"type:varchar(250);not null"`
	Categoria    string  `gorm:"not null"`
	Profesor   string `gorm:"not null"`
}

type Actividades []Actividad