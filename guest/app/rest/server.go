package rest

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	_ "github.com/patricksferraz/guest/app/rest/docs"
	"github.com/patricksferraz/guest/domain/service"
	"github.com/patricksferraz/guest/infra/client/kafka"
	"github.com/patricksferraz/guest/infra/db"
	"github.com/patricksferraz/guest/infra/repo"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

// @title Guest Swagger API
// @version 1.0
// @description Swagger API for Guest Service.
// @termsOfService http://swagger.io/terms/

// @contact.name Coding4u
// @contact.email contato@coding4u.com.br

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
		guest := v1.Group("/guests")
		guest.Post("", restService.CreateGuest)
		guest.Get("/:guest_id", restService.FindGuest)
	}

	addr := fmt.Sprintf("0.0.0.0:%d", port)
	err := r.Listen(addr)
	if err != nil {
		log.Fatal("cannot start rest server", err)
	}

	log.Printf("rest server has been started on port %d", port)
}
