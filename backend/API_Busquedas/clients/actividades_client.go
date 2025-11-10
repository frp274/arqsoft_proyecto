package clients

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
)

type Horario struct {
	Dia        string `json:"dia"`
	HoraInicio string `json:"horaInicio"`
	HoraFin    string `json:"horaFin"`
	Cupo       int    `json:"cupo"`
}

type Actividad struct {
	ID          string    `json:"id"`
	Nombre      string    `json:"nombre"`
	Descripcion string    `json:"descripcion"`
	Profesor    string    `json:"profesor"`
	Tags        []string  `json:"tags"`
	Horarios    []Horario `json:"horario"`
}

var (
	httpClient      *http.Client
	apiActividadesURL string
)

func init() {
	httpClient = &http.Client{
		Timeout: 10 * time.Second,
	}
	
	apiHost := getEnv("API_ACTIVIDADES_HOST", "api_actividades")
	apiPort := getEnv("API_ACTIVIDADES_PORT", "8081")
	apiActividadesURL = fmt.Sprintf("http://%s:%s", apiHost, apiPort)
}

// GetActividadFromAPI fetches an actividad by ID from API_Actividades
func GetActividadFromAPI(id string) (*Actividad, error) {
	url := fmt.Sprintf("%s/actividad/%s", apiActividadesURL, id)
	
	log.Debugf("Fetching actividad from API: %s", url)

	resp, err := httpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to call API_Actividades: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API_Actividades returned status %d: %s", resp.StatusCode, string(body))
	}

	var actividad Actividad
	if err := json.NewDecoder(resp.Body).Decode(&actividad); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	log.Debugf("Successfully fetched actividad %s: %s", id, actividad.Nombre)
	return &actividad, nil
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
