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
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Db *mongo.Database

func GetActividadById(objectID primitive.ObjectID) (model.Actividad, error) {

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
}
/*
func GetActividadesByNombre(nombre string) (model.Actividades, error) {


}
*/


func InsertActividad(actividad model.Actividad) (model.Actividad, error) {
	result := Db.Collection("actividades").FindOne(context.TODO(), bson.M{"nombre": actividad.Nombre})
	if result.Err() == nil {
		log.Errorf("Actividad con nombre '%s' ya existe", actividad.Nombre)
		return model.Actividad{}, fmt.Errorf("actividad con nombre '%s' ya existe", actividad.Nombre)
	}

	_id, err := Db.Collection("actividades").InsertOne(context.TODO(), actividad)
	if err != nil {
		log.Errorf("Error insertando actividad: %v", err)
		return model.Actividad{}, err
	}
	actividad.Id = _id.InsertedID.(primitive.ObjectID)
	return actividad, nil
}


func DeleteActividad(id primitive.ObjectID) error {
	result, err := Db.Collection("actividades").DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		log.Errorf("Error deleting actividad: %v", err)
		return err
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("actividad not found")
	}

	log.Infof("Actividad with id %s deleted", id.Hex())
	return nil
}

func UpdateActividad(a model.Actividad) (model.Actividad, error) {
    filter := bson.M{"_id": a.Id}
    opts := options.Replace().SetUpsert(false)

    res, err := Db.Collection("actividades").ReplaceOne(context.TODO(), filter, a, opts)
    if err != nil {
        return model.Actividad{}, err
    }
    if res.MatchedCount == 0 {
        return model.Actividad{}, fmt.Errorf("actividad not found")
    }
    return a, nil
}
