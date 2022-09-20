package deliveryHttp

import (
	"os"

	"unfinished-api/src/delivery/http/client"
	"unfinished-api/src/delivery/http/store"
	"unfinished-api/src/repository/postgres"
	"unfinished-api/src/repository/s3"
	"unfinished-api/src/usecase"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/gorm"
)

func SetupAndListen(db *gorm.DB, s3Session *session.Session) {
	router := fiber.New()

	router.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	clientRepo := postgres.NewPostgresClientRepository(db)
	storeRepo := postgres.NewPostgresStoreRepository(db)
	imageRepo := s3.NewFilesImageRepository(s3Session)

	clientUsecase := usecase.NewClientUsecase(clientRepo)
	storeUsecase := usecase.NewStoreUsecase(storeRepo, imageRepo)

	client.NewHandler(router, clientUsecase)
	store.NewHandler(router, storeUsecase)

	router.Listen(":" + os.Getenv("API_PORT"))
}
