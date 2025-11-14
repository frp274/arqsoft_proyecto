package actividades

import (
	"encoding/json"
	"fmt"
	"net/http"
	"api_usuarios/dto"
	log "github.com/sirupsen/logrus"
)

func GetActividadById(id int) (dto.ActividadDto, error) {
    var actividad dto.ActividadDto

    url := fmt.Sprintf("%s/actividades/%d", "http://localhost:8081", id)

    response, err := http.Get(url)
    if err != nil {
        log.Errorf("Error calling API_Actividades: %v", err)
        return actividad, fmt.Errorf("error connecting to actividades API: %w", err)
    }
    defer response.Body.Close()

    if response.StatusCode != http.StatusOK {
        return actividad, fmt.Errorf("API_Actividades returned status %d", response.StatusCode)
    }

    if err := json.NewDecoder(response.Body).Decode(&actividad); err != nil {
        return actividad, fmt.Errorf("failed to decode response: %w", err)
    }
	// Ademas debe de eliminar un cupo en la base de datos -----------------------------------------------UNIFICAR ENDPOINT EN UN FUTURO PARA QUE HAGA TODO JUNTO. OSEA QUE BUSQUE LA ACTIVIDAD POR ID Y QUE SAQUE UN CUPO---------------------
	url = fmt.Sprintf("%s/actividades/%d", "http://localhost:8081", id)

	reponse, err := http.Get(url)
	if err != nil {
		log.Errorf("Error calling API_Actividades: %v", err)
		return actividad, fmt.Errorf("error connecting to actividades API: %w", err)
	}
	defer reponse.Body.Close()

    return actividad, nil
}
