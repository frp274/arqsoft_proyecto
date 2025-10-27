package repository

import (
	"context"
	"fmt"
	model "api_actividades/model"
	log "github.com/sirupsen/logrus"
	//"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	//"go.mongodb.org/mongo-driver/mongo/options"
)

var Db *mongo.Database

func GetActividadById(id string) (model.Actividad, error) {

	objectID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		//logeamos la info
		log.Errorf("Error converting id to mongo ID: %v", err)
		return model.Actividad{}, fmt.Errorf("error converting id to mongo ID: %w", err)
	}

	result := Db.Collection("actividades").FindOne(context.TODO(), bson.M{"_id": objectID})

	if result.Err() != nil {
		log.Errorf("Error finding actividad: %v", result.Err())
		return model.Actividad{}, result.Err()
	}

	var actividad model.Actividad

	if err := result.Decode(&actividad); err != nil {
		log.Errorf("Error decoding actividad: %v", err)
		return model.Actividad{}, err
	}

	log.Infof("Actividad encontrada en BD: %v", actividad)
	return actividad, nil
	//result := Db.Preload("Horarios").Where("id = ?", id).First(&actividad)
	//if result.Error != nil {
	//	//return actividad, result.Error
	//}
	//log.Debugf("Act: %v", actividad)
//
	//return actividad
}