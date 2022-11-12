package config

type Config struct {
	Name string
	Port string
}

func NewConfig() *Config {
	return &Config{
		Name: "Max",
		Port: "8000",
	}
}
