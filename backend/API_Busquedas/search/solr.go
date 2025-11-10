package search

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/vanng822/go-solr/solr"
)

var (
	SolrClient *solr.SolrInterface
	SolrCore   = "actividades"
)

// InitSolr initializes the connection to Solr
func InitSolr() error {
	solrHost := getEnv("SOLR_HOST", "solr")
	solrPort := getEnv("SOLR_PORT", "8983")
	solrURL := fmt.Sprintf("http://%s:%s/solr", solrHost, solrPort)

	log.Infof("Connecting to Solr at %s", solrURL)

	si, err := solr.NewSolrInterface(solrURL, SolrCore)
	if err != nil {
		return fmt.Errorf("failed to create Solr interface: %w", err)
	}

	SolrClient = si

	// Test connection
	query := solr.NewQuery()
	query.Q("*:*")
	query.Rows(0)

	searchObj := si.Search(query)
	if _, err := searchObj.Result(nil); err != nil {
		return fmt.Errorf("failed to connect to Solr: %w", err)
	}

	log.Info("Successfully connected to Solr")
	return nil
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
