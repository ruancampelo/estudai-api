package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var DB *gorm.DB

func ConnectDatabase() (*gorm.DB, error) {
	dbConnectionString := os.Getenv("CONNECTION_STRING")

	// Configuração do logger
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // log padrão
		logger.Config{
			SlowThreshold: time.Second, // Definir um limiar para consultas lentas
			LogLevel:      logger.Info, // Nível de log (Info, Warn, Error)
			Colorful:      true,        // Saída colorida
		},
	)

	db, err := gorm.Open(postgres.Open(dbConnectionString), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return nil, fmt.Errorf("falha ao conectar ao banco de dados: %w", err)
	}

	fmt.Println("Conexão com banco de dados estabelecida")
	return db, nil
}
