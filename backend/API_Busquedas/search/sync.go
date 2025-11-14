package search

import (
	"api_busquedas/dto"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

// SyncFromActividades fetches all activities from API_Actividades and indexes them in Solr
func SyncFromActividades() error {
	log.Info("Starting initial sync from API_Actividades to Solr...")

	// Get API_Actividades URL from environment or use default
	apiURL := os.Getenv("API_ACTIVIDADES_URL")
	if apiURL == "" {
		apiURL = "http://api_actividades:8080" // Internal Docker network
	}

	// Give API_Actividades time to start
	time.Sleep(5 * time.Second)

	// Fetch all activities from API_Actividades
	resp, err := http.Get(fmt.Sprintf("%s/actividades", apiURL))
	if err != nil {
		log.Warnf("Could not sync from API_Actividades (might not be ready yet): %v", err)
		return nil // Don't fail startup
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Warnf("API_Actividades returned status %d, skipping sync", resp.StatusCode)
		return nil
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response: %w", err)
	}

	var actividades []dto.ActividadDto
	if err := json.Unmarshal(body, &actividades); err != nil {
		return fmt.Errorf("failed to unmarshal activities: %w", err)
	}

	if len(actividades) == 0 {
		log.Info("No activities to sync")
		return nil
	}

	// Index each activity in Solr
	successCount := 0
	for _, actividad := range actividades {
		if err := indexActivityInSolr(actividad); err != nil {
			log.Errorf("Failed to index activity %s: %v", actividad.Nombre, err)
		} else {
			successCount++
		}
	}

	log.Infof("✅ Initial sync completed: %d/%d activities indexed in Solr", successCount, len(actividades))
	return nil
}

func indexActivityInSolr(actividad dto.ActividadDto) error {
	solrHost := getEnv("SOLR_HOST", "solr")
	solrPort := getEnv("SOLR_PORT", "8983")
	solrURL := fmt.Sprintf("http://%s:%s/solr", solrHost, solrPort)
	
	// Build Solr document
	doc := map[string]interface{}{
		"id":          actividad.Id, // Already a string from MongoDB ObjectID.Hex()
		"nombre":      actividad.Nombre,
		"descripcion": actividad.Descripcion,
		"profesor":    actividad.Profesor,
	}

	// Wrap in Solr update format
	update := []map[string]interface{}{doc}
	payload, err := json.Marshal(update)
	if err != nil {
		return err
	}

	// Send to Solr
	updateURL := fmt.Sprintf("%s/actividades/update?commit=true", solrURL)
	req, err := http.NewRequest("POST", updateURL, strings.NewReader(string(payload)))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("solr returned status %d: %s", resp.StatusCode, string(body))
	}

	return nil
}
