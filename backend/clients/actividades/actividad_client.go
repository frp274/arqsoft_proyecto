package actividad

import (
	"arqsoft_proyecto/model"
	"fmt"
	"strings"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var Db *gorm.DB

func GetActividadById(id int) model.Actividad {
	var actividad model.Actividad

	result := Db.Preload("Horarios").Where("id = ?", id).First(&actividad)
	if result.Error != nil {
		//return actividad, result.Error
	}
	log.Debugf("Act: %v", actividad)

	return actividad
}

func GetAllActividades() (model.Actividades, error) {
	var actividades model.Actividades
	result := Db.Preload("Horarios").Find(&actividades)
	if result.Error != nil {
		return actividades, result.Error
	}

	return actividades, nil
}

/*func GetActividadesFiltradas(nombre string) ([]model.Actividad, error) {
	var actividades []model.Actividad
	query := Db

	if nombre != "" {
		query = query.Where("LOWER(nombre) LIKE ?", "%"+strings.ToLower(nombre)+"%")
	}

	if err := query.Find(&actividades).Error; err != nil {
		return nil, err
	}
	return actividades, nil
}*/

/*func GetActividadesFiltradas(nombre string) ([]model.Actividad, error) {
var actividades []model.Actividad

log.Infof(">> Filtrando actividades con: '%s'", nombre)

query := Db

if nombre != "" {
	// Mostramos cómo se arma la consulta
	/*like := "%" + strings.ToLower(nombre) + "%"
	log.Infof(">> Query LIKE: %s", like)
	query = query.Where("LOWER(nombre) LIKE ?", like)*/

/*query = query.Where("LOWER(nombre) LIKE ?", fmt.Sprintf("%%%s%%", strings.ToLower(nombre)))*/
/*query = query.Where("nombre LIKE ?", "%"+nombre+"%")

	}

	if err := query.Preload("Horarios").Find(&actividades).Error; err != nil {
		log.Errorf("Error en consulta: %v", err)
		return nil, err
	}

	log.Infof(">> Actividades encontradas: %d", len(actividades))

	return actividades, nil
}*/

/*func GetActividadesFiltradas(nombre string) ([]model.Actividad, error) {
	var actividades []model.Actividad

	log.Infof(">> Filtrando actividades con: '%s'", nombre)

	query := Db

	if nombre != "" {
		// ✅ Esta forma es segura y funciona siempre
		query = query.Where("LOWER(nombre) LIKE LOWER(?)", "%"+nombre+"%")
	}

	if err := query.Preload("Horarios").Find(&actividades).Error; err != nil {
		log.Errorf("Error en consulta: %v", err)
		return nil, err
	}

	log.Infof(">> Actividades encontradas: %d", len(actividades))
	return actividades, nil
}*/

func GetActividadesFiltradas(nombre string) ([]model.Actividad, error) {
	var actividades []model.Actividad

	log.Infof(">> Filtro recibido: '%s'", nombre)

	query := Db.Model(&model.Actividad{}).Preload("Horarios")

	if nombre != "" {
		nombreLike := fmt.Sprintf("%%%s%%", strings.ToLower(nombre))
		log.Infof(">> Filtro aplicado (LIKE): %s", nombreLike)
		query = query.Where("LOWER(nombre) LIKE ?", nombreLike)
	}

	if err := query.Find(&actividades).Error; err != nil {
		log.Errorf("Error al buscar actividades: %v", err)
		return nil, err
	}

	log.Infof(">> Actividades encontradas: %d", len(actividades))
	return actividades, nil
}

func InsertActividad(actividad model.Actividad) model.Actividad {
	result := Db.Create(&actividad)

	if result.Error != nil {
		//TODO Manage Errors
		log.Error("")
	}
	log.Debugf("Actividad Created: %v", actividad.Id)
	return actividad
}

func UpdateActividad(actividad model.Actividad) model.Actividad {
	result := Db.Save(&actividad)
	if result.Error != nil {
		log.Error("")
	}
	return actividad
}

func DeleteActividad(id int) error {
	// 1. Borrar los horarios asociados a la actividad
	err := Db.Where("actividad_id = ?", id).Delete(&model.Horario{}).Error
	if err != nil {
		return fmt.Errorf("error al eliminar los horarios asociados: %w", err)
	}

	// 2. Borrar la actividad en sí
	actividad := model.Actividad{}
	result := Db.Delete(&actividad, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no se encontró actividad con ID %d", id)
	}

	return nil
}
