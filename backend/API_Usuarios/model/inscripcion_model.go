package model

type Inscripcion struct {
	Id          int    `gorm:"primaryKey;autoIncrement"`
	UsuarioId   int    `gorm:"not null;index"`
	ActividadId string `gorm:"not null;type:varchar(255);index"`
	HorarioId   string `gorm:"not null;type:varchar(255);index"`
	// Relación con Usuario
	Usuario Usuario `gorm:"foreignKey:UsuarioId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type Inscripciones []Inscripcion
