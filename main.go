package main

import (
	"EffectiveMobile/models"
	"github.com/SZabrodskii/microservices/libs/providers"
	"github.com/SZabrodskii/microservices/libs/providers/config"
	"github.com/SZabrodskii/microservices/libs/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"log"
)

func main() {

	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Printf("Failed to initialize logger: %v", err)
	}
	defer logger.Sync()

	cfg := &config.PostgreSQLParams{
		Host:     utils.GetEnv("DB_HOST", "localhost"),
		Port:     utils.GetEnv("DB_PORT", 5432),
		Username: utils.GetEnv("DB_USERNAME", ""),
		Password: utils.GetEnv("DB_PASSWORD", ""),
		Database: utils.GetEnv("DB_NAME", ""),
		TLS:      nil,
	}
	postgreSQLProvider, err := providers.NewPostgreSQLProvider(cfg, logger)
	if err != nil {
		logger.Error("Failed to initialize PostgreSQL provider", zap.Error(err))
	}

	db := postgreSQLProvider.DB()

	r := Repository{
		DB: db,
	}

	err = db.AutoMigrate(&models.Car{}, &models.People{})
	if err != nil {
		logger.Error("Could not migrate DB: %v", zap.Error(err))
	}

	router := gin.Default()
	r.SetupRoutes(router)
	router.Run(":5432")
}
