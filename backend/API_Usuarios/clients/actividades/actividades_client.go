package actividades

import (
	"api_usuarios/dto"
	"encoding/json"
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func GetActividadById(id string) (dto.ActividadDto, error) {
	var actividad dto.ActividadDto

	url := fmt.Sprintf("%s/actividad/%s", "http://localhost:8081", id)

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
	url = fmt.Sprintf("%s/actividad/%s/borar-cupo", "http://localhost:8081", id)

	reponse, err := http.Post(url, "application/json", nil)
	if err != nil {
		log.Errorf("Error calling API_Actividades para borrar cupo: %v", err)
		return actividad, fmt.Errorf("error connecting to actividades API for quota: %w", err)
	}
	defer reponse.Body.Close()

	return actividad, nil
}
