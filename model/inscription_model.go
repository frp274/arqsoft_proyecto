package model

type Inscripcion struct {
	Id          int    `gorm:"primaryKey"`
	Fecha       string `gorm:"not null"`
	Usuario     Usuario   `gorm:"foreignKey:UserId"`
	UsuarioId   int
	Actividad   Actividad `gorm:"foreignKey:ActivityId"`
	ActividadId int
}

type Inscripciones []Inscripcion
