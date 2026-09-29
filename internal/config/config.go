package config

import "os"

type Config struct {
	Host       string
	User       string
	Password   string
	Dbname     string
	Port       string
	Seed       string
	ServerPort string
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

func CreateConfig() Config {
	config := Config{
		Host:       getEnv("DBHOST", "localhost"),
		User:       getEnv("DBUSER", "root"),
		Password:   getEnv("DBPASSWORD", "password"),
		Dbname:     getEnv("DBNAME", "db"),
		Port:       getEnv("DBPORT", "5432"),
		Seed:       getEnv("SEEDDB", "true"),
		ServerPort: getEnv("SERVERPORT", "1323"),
	}
	return config
}
