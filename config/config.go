package config

type Configuration struct {
	DatabaseDSN string `env:"DATABASE_DSN"`
}

var ApplicationConfig = Configuration{}
