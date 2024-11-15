package main

import (
	. "estudai-api/api"
	. "estudai-api/internal/infrastructure/database"
	. "estudai-api/internal/infrastructure/dependency"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	modeEnv := os.Getenv("MODE_ENVIRONMENT")

	if modeEnv == "release" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		err := godotenv.Load()
		if err != nil {
			log.Fatalf("Error loading .env file")
		}
	}
	router := gin.Default()

	dbInstance, err := ConnectDatabase()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	deps := InitDependencies(dbInstance)
	RegisterRoutes(router, deps)

	certFile := os.Getenv("CERTIFICATE")
	keyFile := os.Getenv("KEY_CERTIFCATE")

	if err := router.RunTLS(":5112", certFile, keyFile); err != nil {
		log.Fatalf("Erro ao iniciar o servidor HTTPS: %s", err)
	}
	router.Run(":5112")

}
