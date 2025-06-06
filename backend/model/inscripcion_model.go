package model

type Inscripcion struct {
	Id          int    `gorm:"primaryKey"`
	Fecha       string `gorm:"not null"`
	Usuario     Usuario   `gorm:"foreignKey:UsuarioId"`
	UsuarioId   int
	Actividad   Actividad `gorm:"foreignKey:ActividadId"`
	ActividadId int
}

type Inscripciones []Inscripcion
