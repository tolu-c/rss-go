// Package environment loads and validates configuration from the process
// environment. Exposing a single Config struct (rather than os.Getenv calls
// scattered across the codebase) keeps required env vars discoverable and
// makes testing easier — tests can build a Config directly.
package environment

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Config is the typed view of process environment variables consumed by the
// application. Adding a new env var means adding a field here and a check in
// Load — there is no other place env vars should be read.
type Config struct {
	Port  string
	DBURL string
}

// Load reads .env (best-effort) and pulls the required variables out of the
// process environment. It returns an error listing any missing required
// values rather than calling log.Fatal so callers can decide how to handle
// startup failures.
func Load() (*Config, error) {
	// godotenv.Load is best-effort: missing .env in production is normal.
	_ = godotenv.Load(".env")

	cfg := &Config{
		Port:  os.Getenv("PORT"),
		DBURL: os.Getenv("DB_URL"),
	}

	var missing []string
	if cfg.Port == "" {
		missing = append(missing, "PORT")
	}
	if cfg.DBURL == "" {
		missing = append(missing, "DB_URL")
	}
	if len(missing) > 0 {
		return nil, fmt.Errorf("environment: missing required vars: %v", missing)
	}

	return cfg, nil
}
