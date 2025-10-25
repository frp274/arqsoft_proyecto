package repository

import (
	"context"
	"fmt"
	model "api_actividades/model"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


type Actividad struct {
	Id          int      
	Nombre      string   
	Descripcion string   
	Profesor    string   
	Horarios    []Horario
}

type Horario struct {
	Id           int	
	ActividadID  int
	Actividad    Actividad
	Dia          string   
	HoraInicio   string   
	HoraFin      string   
	Cupo         int      
}

type Actividades []Actividad


func GetActividadById(id int) (model.Actividad, error) {
	var actividad model.Actividad

	objectID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return model.Actividad{}, fmt.Errorf("error converting id to mongo ID: %w", err)
	}

	result := Db.Collection("actividades").FindOne(context.TODO(), bson.M{"_id": objectID})

	//result := Db.Preload("Horarios").Where("id = ?", id).First(&actividad)
	//if result.Error != nil {
	//	//return actividad, result.Error
	//}
	//log.Debugf("Act: %v", actividad)
//
	//return actividad
}