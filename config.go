package main

type Config struct {
	ApiKey string `envconfig:"API_KEY" required:"true"`
	Ports  struct {
		GRAPHQL string `envconfig:"GRAPHQL_PORT" default:"8081"`
	}
}

func (c Config) Validate() error {
	return nil
}
