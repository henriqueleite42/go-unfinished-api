package main

import (
	"os"
	deliveryHttp "unfinished-api/src/delivery/http"
	"unfinished-api/src/repository/postgres"
	"unfinished-api/src/repository/s3"

	"github.com/joho/godotenv"
)

func main() {
	if os.Getenv("ENV") == "dev" {
		godotenv.Load(".env.dev")
	} else {
		godotenv.Load(".env")
	}

	postgresSQL := postgres.GetPostgres()
	s3Session := s3.GetS3()

	postgres.RunMigrations(postgresSQL)

	deliveryHttp.SetupAndListen(postgresSQL, s3Session)
}
