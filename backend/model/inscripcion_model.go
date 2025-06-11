package model

type Inscripcion struct {
	Id          		int    		`gorm:"primaryKey"`
	Horario				Horario		`gorm:"foreignKey:HorarioId"`
	HorarioId  			int 
	Usuario     		Usuario  	`gorm:"foreignKey:UsuarioId"`
	UsuarioId   		int
	Actividad   		Actividad 	`gorm:"foreignKey:ActividadId"`
	ActividadId 		int
}

type Inscripciones []Inscripcion
