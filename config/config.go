package config

type AppConfig struct {
    Port int
    // Add more configuration options here
}

func LoadConfig() AppConfig {
    // Load configuration from file or environment variables
    return AppConfig{Port: 8080}
}
