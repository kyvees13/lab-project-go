package config

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName    string
	Debug      bool
	DBConnStr  string
	ServerPort int
}

const projectDirName = "lab-server"

func LoadConfig() *Config {

	projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))
	err := godotenv.Load(string(rootPath) + `/.env`)
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	time.Local = time.FixedZone("MSK", 3*60*60) // Устанавливаем часовой пояс

	return &Config{
		AppName:    getEnv("APP_NAME", "Default project"),
		Debug:      getEnvAsBool("DEBUG", true),
		DBConnStr:  getEnv("DB_CONN", ""),
		ServerPort: getEnvAsInt("SERVER_PORT", 8080),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	valueStr := getEnv(key, "")
	if _, err := strconv.Atoi(valueStr); err != nil {
		log.Fatalf("Error converting %s to int", valueStr)
	}
	return defaultValue
}

func getEnvAsBool(key string, defaultValue bool) bool {
	valueStr := getEnv(key, "")
	if value, err := strconv.ParseBool(valueStr); err == nil {
		return value
	}
	return defaultValue
}
