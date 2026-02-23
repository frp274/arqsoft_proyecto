package search

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
)

var (
	SolrClient *SolrInterface
	SolrCore   = "actividades"
)

type SolrInterface struct {
	client  *http.Client
	baseURL string
}

// InitSolr initializes the connection to Solr
func InitSolr() error {
	solrHost := getEnv("SOLR_HOST", "solr")
	solrPort := getEnv("SOLR_PORT", "8983")
	solrURL := fmt.Sprintf("http://%s:%s/solr/%s", solrHost, solrPort, SolrCore)

	log.Infof("Connecting to Solr at %s", solrURL)

	SolrClient = &SolrInterface{
		client:  &http.Client{Timeout: 10 * time.Second},
		baseURL: solrURL,
	}

	// Test connection using ping
	resp, err := SolrClient.client.Get(solrURL + "/admin/ping")
	if err != nil {
		return fmt.Errorf("failed to ping Solr: %w", err)
	}
	defer resp.Body.Close()
	
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("solr ping returned status %d", resp.StatusCode)
	}

	log.Info("Successfully connected to Solr using net/http")
	return nil
}

// Update sends a document to index
func (s *SolrInterface) Update(doc map[string]interface{}) error {
	payload := []map[string]interface{}{doc}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	
	req, _ := http.NewRequest("POST", s.baseURL+"/update/json/docs?commit=true", bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")
	
	resp, err := s.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	
	if resp.StatusCode >= 400 {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("solr update returned status %d: %s", resp.StatusCode, string(body))
	}
	return nil
}

func (s *SolrInterface) Delete(id string) error {
	payload := map[string]interface{}{
		"delete": map[string]interface{}{"id": id},
	}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	
	req, _ := http.NewRequest("POST", s.baseURL+"/update?commit=true", bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")
	
	resp, err := s.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	
	if resp.StatusCode >= 400 {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("solr delete returned status %d: %s", resp.StatusCode, string(body))
	}
	return nil
}

type SearchResult struct {
	Response struct {
		NumFound int                      `json:"numFound"`
		Docs     []map[string]interface{} `json:"docs"`
	} `json:"response"`
}

func (s *SolrInterface) Search(query, sort string, start, rows int) (*SearchResult, error) {
	u, _ := url.Parse(s.baseURL + "/select")
	qVar := u.Query()
	qVar.Set("q", query)
	qVar.Set("start", fmt.Sprintf("%d", start))
	qVar.Set("rows", fmt.Sprintf("%d", rows))
	if sort != "" {
		qVar.Set("sort", sort)
	}
	u.RawQuery = qVar.Encode()
	
	resp, err := s.client.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	
	if resp.StatusCode >= 400 {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("solr search returned status %d: %s", resp.StatusCode, string(body))
	}
	
	var result SearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
