package model

type Inscripcion struct {
	Id int `gorm:"primaryKey"`
	// Horario y Actividad sin constraint para permitir IDs que no existen en MySQL
	// Los datos maestros están en MongoDB
	HorarioId   string `gorm:"index"`
	UsuarioId   int    `gorm:"index"`
	ActividadId string `gorm:"index"`
	// Relaciones opcionales (solo funcionan si los datos están en MySQL)
	Usuario *Usuario `gorm:"foreignKey:UsuarioId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

type Inscripciones []Inscripcion
