package rest

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	_ "github.com/patricksferraz/menu/app/rest/docs"
	"github.com/patricksferraz/menu/domain/service"
	"github.com/patricksferraz/menu/infra/client/kafka"
	"github.com/patricksferraz/menu/infra/db"
	"github.com/patricksferraz/menu/infra/repo"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

// @title Menu Swagger API
// @version 1.0
// @description Swagger API for Menu Service.
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
		menu := v1.Group("/menus")
		menu.Post("", restService.CreateMenu)
		menu.Get("/:menu_id", restService.FindMenu)

		item := menu.Group("/:menu_id/items")
		item.Post("", restService.CreateItem)
		item.Get("/:item_id", restService.FindItem)
		item.Put("/:item_id/tag", restService.AddItemTag)
	}

	addr := fmt.Sprintf("0.0.0.0:%d", port)
	err := r.Listen(addr)
	if err != nil {
		log.Fatal("cannot start rest server", err)
	}

	log.Printf("rest server has been started on port %d", port)
}
