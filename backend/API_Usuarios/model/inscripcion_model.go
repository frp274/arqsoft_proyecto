package model

type Inscripcion struct {
	Id          int    `gorm:"primaryKey;autoIncrement"`
	UsuarioId   int    `gorm:"not null;index"`
	ActividadId int    `gorm:"not null;index"`
	HorarioId   int    `gorm:"not null;index"`
	// Relación con Usuario
	Usuario     Usuario `gorm:"foreignKey:UsuarioId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type Inscripciones []Inscripcion
