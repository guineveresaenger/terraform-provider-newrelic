package config

import (
	"net/http"
	"time"

	"github.com/newrelic/newrelic-client-go/internal/logging"
	"github.com/newrelic/newrelic-client-go/internal/version"
)

// RegionType represents the members of the Region enumeration.
type RegionType int

// DefaultBaseURLs represents the base API URLs for the different environments of the New Relic REST API V2.
var DefaultBaseURLs = map[RegionType]string{
	Region.US:      "https://api.newrelic.com/v2",
	Region.EU:      "https://api.eu.newrelic.com/v2",
	Region.Staging: "https://staging-api.newrelic.com/v2",
}

const (
	// US represents New Relic's US-based production deployment.
	US = iota

	// EU represents New Relic's EU-based production deployment.
	EU

	// Staging represents New Relic's US-based staging deployment.
	// This is for internal New Relic use only.
	Staging
)

// Region specifies the New Relic environment to target.
var Region = struct {
	US      RegionType
	EU      RegionType
	Staging RegionType
}{
	US:      US,
	EU:      EU,
	Staging: Staging,
}

// Config contains all the configuration data for the API Client.
type Config struct {
	BaseURL       string
	APIKey        string
	Timeout       *time.Duration
	HTTPTransport *http.RoundTripper
	UserAgent     string
	Region        RegionType

	// LogLevel can be one of the following values:
	// "panic", "fatal", "error", "warn", "info", "debug", "trace"
	LogLevel string
	LogJSON  bool
	Logger   logging.Logger
}

// GetLogger returns a logger instance based on the config values.
func (c *Config) GetLogger() logging.Logger {
	if c.Logger != nil {
		return c.Logger
	} else {
		l := logging.NewStructuredLogger().
			SetDefaultFields(map[string]string{"newrelic-client-go": version.Version}).
			LogJSON(c.LogJSON).
			SetLogLevel(c.LogLevel)

		c.Logger = l
		return l
	}
}
