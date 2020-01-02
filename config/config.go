package config

type config struct {
	DatabaseDSN string `env:"DATABASE_DSN"`
}

// ApplicationConfig configuration of the app
var ApplicationConfig = config{}
