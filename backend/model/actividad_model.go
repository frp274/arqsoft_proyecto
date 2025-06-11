package model

type Actividad struct {
	Id          int    `gorm:"primaryKey"`
	Nombre      string `gorm:"not null"`
	Descripcion string `gorm:"type:varchar(250);not null"`
	Profesor    string `gorm:"not null"`
	Cupo        int    `gorm:"not null"`
	Horario     []Horario `gorm:"not null"`
}

type Horario struct {
	Dia        string
	HoraInicio string
	HoraFin    string
}

type Actividades []Actividad
