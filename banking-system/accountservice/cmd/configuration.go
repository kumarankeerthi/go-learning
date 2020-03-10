package cmd

type Config struct {
	AppName            string
	CockroachdbConnURL string
	ServicePort        string
}

func DefaultConfig() *Config {
	return &Config{
		AppName:            "AccountService",
		CockroachdbConnURL: "postgresql://keerthi@localhost:26257/bank?sslmode=disable",
		ServicePort:        "8501",
	}
}
