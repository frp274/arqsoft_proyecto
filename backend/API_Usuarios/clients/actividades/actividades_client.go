package actividades

import (
	"api_usuarios/dto"
	"encoding/json"
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func GetActividadById(id string, horarioId string) (dto.ActividadDto, error) {
	var actividad dto.ActividadDto

	// Use service name instead of localhost for Docker networking
	baseURL := "http://api_actividades:8080"
	url := fmt.Sprintf("%s/actividad/%s", baseURL, id)

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

	// Incrementally update the cupo in API_Actividades
	url = fmt.Sprintf("%s/actividad/%s/borar-cupo?horario_id=%s", baseURL, id, horarioId)

	respCupo, err := http.Post(url, "application/json", nil)
	if err != nil {
		log.Errorf("Error calling API_Actividades para borrar cupo: %v", err)
		return actividad, fmt.Errorf("error connecting to actividades API for quota: %w", err)
	}
	defer respCupo.Body.Close()

	if respCupo.StatusCode != http.StatusOK {
		log.Errorf("API_Actividades /borar-cupo returned status %d", respCupo.StatusCode)
	}

	return actividad, nil
}
