package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func LoadEnv(fileConfig string) (map[string]string, error) {
	webConfig := map[string]string{}
	err := godotenv.Load(fileConfig)
	if err != nil {
		log.Error("Error loading .env file")
		return nil, errors.New("error loading .env file")
	}
	webConfig["jwtKey"] = os.Getenv("JWT_SECRET")
	webConfig["appName"] = os.Getenv("APP_NAME")
	webConfig["appVersion"] = os.Getenv("APP_VERSION")
	webConfig["port"] = os.Getenv("PORT")
	webConfig["grpcHost"] = os.Getenv("GRPC_HOST")
	return webConfig, nil
}
