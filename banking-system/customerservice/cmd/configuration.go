// Package cmd is the entry point this service
package cmd

// Config type will hold all the config strins needed for this service
type Config struct {
	Environment        string
	CockroachdbConnURL string
	ServicePort        string
}

// DefaultConfig will return the config type with detauld values
func DefaultConfig() *Config {
	return &Config{
		Environment:        "local",
		CockroachdbConnURL: "postgresql://keerthi@localhost:26257/bank?sslmode=disable",
		ServicePort:        "8500",
	}
}
