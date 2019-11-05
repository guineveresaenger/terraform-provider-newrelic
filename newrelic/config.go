package newrelic

import (
	"log"

	"github.com/hashicorp/terraform/helper/logging"

	synthetics "github.com/dollarshaveclub/new-relic-synthetics-go"
	newrelic "github.com/paultyng/go-newrelic/v4/api"
)

// Config contains New Relic provider settings
type Config struct {
	APIKey    string
	APIURL    string
	userAgent string
}

// Client returns a new client for accessing New Relic
func (c *Config) Client() (*newrelic.Client, error) {
	nrConfig := newrelic.Config{
		APIKey:    c.APIKey,
		BaseURL:   c.APIURL,
		Debug:     logging.IsDebugOrHigher(),
		UserAgent: c.userAgent,
	}

	client := newrelic.New(nrConfig)

	log.Printf("[INFO] New Relic client configured")

	return &client, nil
}

// ClientInfra returns a new client for accessing New Relic
func (c *Config) ClientInfra() (*newrelic.InfraClient, error) {
	nrConfig := newrelic.Config{
		APIKey:    c.APIKey,
		BaseURL:   c.APIURL,
		Debug:     logging.IsDebugOrHigher(),
		UserAgent: c.userAgent,
	}

	client := newrelic.NewInfraClient(nrConfig)

	log.Printf("[INFO] New Relic Infra client configured")

	return &client, nil
}

// ClientSynthetics returns a new client for accessing New Relic Synthetics
func (c *Config) ClientSynthetics() (*synthetics.Client, error) {
	conf := func(s *synthetics.Client) {
		s.APIKey = c.APIKey
	}

	client, _ := synthetics.NewClient(conf)

	log.Printf("[INFO] New Relic Synthetics client configured")

	return client, nil
}

// ProviderConfig for the custom provider
type ProviderConfig struct {
	Client      *newrelic.Client
	InfraClient *newrelic.InfraClient
	Synthetics  *synthetics.Client
}
