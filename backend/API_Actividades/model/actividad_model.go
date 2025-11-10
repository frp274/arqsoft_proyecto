package model

import "go.mongodb.org/mongo-driver/bson/primitive"
/*

// Estructura vieja con GORM para MySQL
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
    HoraInicio   string    `gorm:"not null"`
    HoraFin      string    `gorm:"not null"`
    Cupo         int       `gorm:"not null"`
}

//Primera prueba de struct con MongoDB
type Actividad struct {
    ID          primitive.ObjectID `bson:"_id,omitempty"`
    Nombre      string             `bson:"nombre"`
    Descripcion string             `bson:"descripcion,omitempty"`
    //CreatedAt   time.Time          `bson:"createdAt"`
    //UpdatedAt   time.Time          `bson:"updatedAt"`
}

type Horario struct {
    ID          primitive.ObjectID `bson:"_id,omitempty"`
    Dia   int                `bson:"dia"`   // 0..6
    HoraInicio  string             `bson:"horaInicio"`  // "HH:mm"
    HoraFin     string             `bson:"horaFin"`
    Cupo        int                `bson:"cupo"`
    Disponibles int                `bson:"disponibles"`
    Estado      string             `bson:"estado"`      // activo/inactivo
    //CreatedAt   time.Time          `bson:"createdAt"`
    //UpdatedAt   time.Time          `bson:"updatedAt"`
}
*/

// Horario define la estructura de cada elemento dentro del array 'horarios'.
type Horario struct {
	Dia  string `bson:"dia" json:"dia"`
	HoraInicio string `bson:"horaInicio" json:"horaInicio"`
	HoraFin    string `bson:"horaFin" json:"horaFin"`
	Cupo       int    `bson:"cupo" json:"cupo"` // 'int' porque el schema dice "int" y tiene min/max.
}

// Actividad representa el documento principal que se almacenará en la colección.
type Actividad struct {
	// ID es el identificador único de MongoDB.
	Id primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Nombre string `bson:"nombre" json:"nombre"`// 'nombre' está requerido y es un string.
	Profesor string `bson:"profesor" json:"profesor"`// 'profesor' está requerido y es un string.
	OwnerId int `bson:"owner_id" json:"owner_id"` // ID del usuario creador/dueño (validado contra API_Usuarios)
	Horarios []Horario `bson:"horarios" json:"horarios"`// 'horarios' está requerido y es un array de objetos 'Horario'.
	Descripcion string `bson:"descripcion,omitempty" json:"descripcion,omitempty"`// 'descripcion' es un array de strings o null, lo más simple es usar []string o un puntero. El schema dice: bson.A{"string", "null"} Si solo quieres strings, usa []string. Si necesitas el manejo de nulos:
	Tags []string `bson:"tags,omitempty" json:"tags,omitempty"`// 'tags' es un array de strings opcional.

	// Campos de auditoría (asumiendo que los vas a añadir al final, aunque no estén en el schema)
	//CreatedAt primitive.DateTime `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
	//UpdatedAt primitive.DateTime `bson:"updatedAt,omitempty" json:"updatedAt,omitempty"`
}

type Actividades []Actividad
