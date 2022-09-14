package main

type Config struct {
	ApiKey string `envconfig:"API_KEY" required:"true" default:"b5c572d75e504a1282f5e10f7f5bceb5"`
	Ports  struct {
		GRAPHQL string `envconfig:"GRAPHQL_PORT" default:"8081"`
	}
}

func (c Config) Validate() error {
	return nil
}
