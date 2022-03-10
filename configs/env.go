package configs

import (
	"gin-mongo-api/logger"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func loadEnv() {
	if err := godotenv.Load(); err != nil {
		logger.ErrorLogger.Fatalln("Error loading .env file")
	}
}

func EnvMongoURI() string {
	loadEnv()
	return os.Getenv("MONGOURI")
}

func EnvDatabaseName() string {
	loadEnv()
	return os.Getenv("DATABASE")
}

func EnvPageLimit() int64 {
	loadEnv()
	pageLimit, _ := strconv.ParseInt(os.Getenv("PAGE_LIMIT"), 10, 64)
	return pageLimit
}
