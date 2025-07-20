package model

type Actividad struct {
	Id          int       `gorm:"primaryKey"`
	Nombre      string    `gorm:"not null"`
	Descripcion string    `gorm:"type:varchar(250);not null"`
	Profesor    string    `gorm:"not null"`
	Horarios    []Horario `gorm:"foreignKey:ActividadID"`
}

type Horario struct {
    Id           int		`gorm:"primaryKey"`
    ActividadID  int
    Actividad    Actividad `gorm:"foreignKey:ActividadID"`
    Dia          string    `gorm:"not null"`
    HoraInicio   string
    HoraFin      string    `gorm:"not null"`
    Cupo         int       `gorm:"not null"`
}


type Actividades []Actividad

