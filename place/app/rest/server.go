package rest

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	_ "github.com/patricksferraz/place/app/rest/docs"
	"github.com/patricksferraz/place/domain/service"
	"github.com/patricksferraz/place/infra/client/kafka"
	"github.com/patricksferraz/place/infra/db"
	"github.com/patricksferraz/place/infra/repo"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

// @title Place Swagger API
// @version 1.0
// @description Swagger API for Place Service.
// @termsOfService http://swagger.io/terms/

// @contact.name Patrick Ferraz
// @contact.email patrick.ferraz@outlook.com

// @BasePath /api/v1
// @query.collection.format multi

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func StartRestServer(pg *db.PostgreSQL, kp *kafka.KafkaProducer, port int) {
	r := fiber.New()
	r.Use(cors.New())

	repository := repo.NewRepository(pg, kp)
	service := service.NewService(repository)
	restService := NewRestService(service)

	api := r.Group("/api")

	v1 := api.Group("/v1")
	v1.Get("/swagger/*", fiberSwagger.WrapHandler)
	{
		place := v1.Group("/places")
		place.Post("", restService.CreatePlace)
		place.Get("/:place_id", restService.FindPlace)
	}

	addr := fmt.Sprintf("0.0.0.0:%d", port)
	err := r.Listen(addr)
	if err != nil {
		log.Fatal("cannot start rest server", err)
	}

	log.Printf("rest server has been started on port %d", port)
}
