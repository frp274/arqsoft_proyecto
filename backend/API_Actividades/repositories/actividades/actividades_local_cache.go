package repository

import (
	"api_actividades/model"
	"fmt"
	"time"

	"github.com/karlseguin/ccache"
	log "github.com/sirupsen/logrus"
)

var client *ccache.Cache
var duration time.Duration

type CacheConfig struct {
	MaxSize      int64
	ItemsToPrune uint32
	Duration     time.Duration
}

func NewCache(config CacheConfig) {
	client =ccache.New(ccache.Configure().
		MaxSize(config.MaxSize).
		ItemsToPrune(config.ItemsToPrune))
	duration = config.Duration
	//return Cache{
	//	client:   client,
	//	duration: config.Duration,
	//}
}

func GetActividadByIdCache(id string) (model.Actividad, error) {

	item := client.Get(id)
	if item == nil {
		err := fmt.Errorf("actividad con ID %s no encontrada en local cache", id)
		return model.Actividad{}, err
	}

	if item.Expired() {
		err :=fmt.Errorf("actividad con ID %s expirada en local cache", id)
		return model.Actividad{}, err
	}

	actividad, ok := item.Value().(model.Actividad)
	if !ok {
		err := fmt.Errorf("error al convertir el valor almacenado en local cache para ID %s", id)
		return model.Actividad{}, err
	}

	log.Infof("Actividad encontrada en local cache: %v", actividad)
	return actividad, nil
}

func InsertActividadCache(actividad model.Actividad) model.Actividad {
	log.Infof("Buscando actividad en local cache. Actividad: %v", actividad)
	client.Set(actividad.Id.Hex(), actividad, duration)

	return actividad
}

func DeleteActividadCache(id string) error {
	log.Infof("Eliminando actividad de la local cache. ID: %s", id)

	bool := client.Delete(id)
	if !bool {
		err := fmt.Errorf("actividad con ID %s no encontrada en local cache para eliminar", id)
		return err
	}
	return nil
}
