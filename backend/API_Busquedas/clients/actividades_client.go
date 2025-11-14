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

// HorarioAPI es el formato que viene de API_Actividades
type HorarioAPI struct {
	Dia            string `json:"dia"`
	HorarioInicio  string `json:"horarioInicio"`
	HorarioFinal   string `json:"horarioFinal"`
	Cupo           int    `json:"cupo"`
}

// ActividadAPI es el formato que viene de API_Actividades
type ActividadAPI struct {
	ID          string       `json:"id"`
	Nombre      string       `json:"nombre"`
	Descripcion string       `json:"descripcion"`
	Profesor    string       `json:"profesor"`
	Tags        []string     `json:"tags"`
	Horarios    []HorarioAPI `json:"horarios"`
}

// Horario es el formato que exponemos al frontend
type Horario struct {
	ID            string `json:"id,omitempty"`
	Dia           string `json:"dia"`
	HorarioInicio string `json:"horarioInicio"`
	HorarioFinal  string `json:"horarioFinal"`
	Cupo          int    `json:"cupo"`
}

// Actividad es el formato que exponemos al frontend
type Actividad struct {
	ID          string    `json:"id"`
	Nombre      string    `json:"nombre"`
	Descripcion string    `json:"descripcion"`
	Profesor    string    `json:"profesor"`
	Tags        []string  `json:"tags"`
	Horarios    []Horario `json:"horarios"`
}

var (
	httpClient        *http.Client
	apiActividadesURL string
)

func init() {
	httpClient = &http.Client{
		Timeout: 10 * time.Second,
	}

	apiHost := getEnv("API_ACTIVIDADES_HOST", "api_actividades")
	apiPort := getEnv("API_ACTIVIDADES_PORT", "8080")
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

	// Decodificar en el formato de la API
	var actividadAPI ActividadAPI
	if err := json.NewDecoder(resp.Body).Decode(&actividadAPI); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	// Mapear a nuestro formato
	actividad := &Actividad{
		ID:          actividadAPI.ID,
		Nombre:      actividadAPI.Nombre,
		Descripcion: actividadAPI.Descripcion,
		Profesor:    actividadAPI.Profesor,
		Tags:        actividadAPI.Tags,
		Horarios:    make([]Horario, len(actividadAPI.Horarios)),
	}

	// Mapear horarios con IDs generados
	for i, h := range actividadAPI.Horarios {
		actividad.Horarios[i] = Horario{
			ID:            fmt.Sprintf("%d", i),
			Dia:           h.Dia,
			HorarioInicio: h.HorarioInicio,
			HorarioFinal:  h.HorarioFinal,
			Cupo:          h.Cupo,
		}
	}

	log.Debugf("Successfully fetched actividad %s: %s", id, actividad.Nombre)
	return actividad, nil
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
