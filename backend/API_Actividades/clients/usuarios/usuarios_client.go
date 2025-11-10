package usuarios

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
)

type UsuarioResponse struct {
	Id              int    `json:"id"`
	Nombre_apellido string `json:"nombre_apellido"`
	UserName        string `json:"username"`
	Es_admin        bool   `json:"es_admin"`
}

var (
	usuariosAPIURL string
	httpClient     *http.Client
)

func init() {
	// Leer URL de API_Usuarios desde variable de entorno
	usuariosAPIURL = os.Getenv("USUARIOS_API_URL")
	if usuariosAPIURL == "" {
		usuariosAPIURL = "http://localhost:8082" // Default para desarrollo
	}

	// Cliente HTTP con timeout
	httpClient = &http.Client{
		Timeout: 5 * time.Second,
	}
}

// ValidateUser verifica que un usuario existe en API_Usuarios
func ValidateUser(userId int) error {
	url := fmt.Sprintf("%s/usuario/%d", usuariosAPIURL, userId)

	log.Debugf("Validating user %d against %s", userId, url)

	// Nota: En producción, aquí deberías incluir un token de servicio
	// para autenticación entre microservicios
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}

	// TODO: Agregar token de servicio para autenticación entre microservicios
	// req.Header.Set("Authorization", "Bearer "+serviceToken)

	resp, err := httpClient.Do(req)
	if err != nil {
		log.Errorf("Error calling API_Usuarios: %v", err)
		return fmt.Errorf("error connecting to users API: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		log.Warnf("User %d not found in API_Usuarios", userId)
		return fmt.Errorf("user with ID %d not found", userId)
	}

	if resp.StatusCode != http.StatusOK {
		log.Errorf("API_Usuarios returned status %d for user %d", resp.StatusCode, userId)
		return fmt.Errorf("error validating user: status %d", resp.StatusCode)
	}

	var usuario UsuarioResponse
	if err := json.NewDecoder(resp.Body).Decode(&usuario); err != nil {
		log.Errorf("Error decoding user response: %v", err)
		return fmt.Errorf("error decoding user response: %w", err)
	}

	log.Infof("User %d validated successfully: %s", userId, usuario.UserName)
	return nil
}
