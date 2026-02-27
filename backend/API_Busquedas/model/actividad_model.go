package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// Horario define la estructura de cada elemento dentro del array 'horarios'.
type Horario struct {
	Dia        string `bson:"dia" json:"dia"`
	HoraInicio string `bson:"horaInicio" json:"horaInicio"`
	HoraFin    string `bson:"horaFin" json:"horaFin"`
	Cupo       int    `bson:"cupo" json:"cupo"` // 'int' porque el schema dice "int" y tiene min/max.
}

// Actividad representa el documento principal
type Actividad struct {
	Id          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Nombre      string             `bson:"nombre" json:"nombre"`
	Profesor    string             `bson:"profesor" json:"profesor"`
	OwnerId     int                `bson:"owner_id" json:"owner_id"`
	Horarios    []Horario          `bson:"horarios" json:"horarios"`
	Descripcion string             `bson:"descripcion,omitempty" json:"descripcion,omitempty"`
	ImagenURL   string             `bson:"imagen_url,omitempty" json:"imagen_url,omitempty"`
	Tags        []string           `bson:"tags,omitempty" json:"tags,omitempty"`
}

type Actividades []Actividad
