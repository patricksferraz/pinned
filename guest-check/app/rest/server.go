package rest

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	_ "github.com/patricksferraz/guest-check/app/rest/docs"
	"github.com/patricksferraz/guest-check/domain/service"
	"github.com/patricksferraz/guest-check/infra/client/kafka"
	"github.com/patricksferraz/guest-check/infra/db"
	"github.com/patricksferraz/guest-check/infra/repo"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

// @title Guest Check Swagger API
// @version 1.0
// @description Swagger API for Guest Check Service.
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
		guestCheck := v1.Group("/guest-checks")
		guestCheck.Post("", restService.CreateGuestCheck)
		guestCheck.Get("/:guest_check_id", restService.FindGuestCheck)
		guestCheck.Post("/:guest_check_id/wait-payment", restService.WaitPaymentGuestCheck)
		guestCheck.Post("/:guest_check_id/cancel", restService.CancelGuestCheck)
		guestCheck.Post("/:guest_check_id/pay", restService.PayGuestCheck)

		guestCheckItem := guestCheck.Group("/:guest_check_id/items")
		guestCheckItem.Post("", restService.AddGuestCheckItem)
		guestCheckItem.Get("/:guest_check_item_id", restService.FindGuestCheckItem)
		guestCheckItem.Post("/:guest_check_item_id/cancel", restService.CancelGuestCheckItem)
		guestCheckItem.Post("/:guest_check_item_id/prepare", restService.PrepareGuestCheckItem)
		guestCheckItem.Post("/:guest_check_item_id/ready", restService.ReadyGuestCheckItem)
		guestCheckItem.Post("/:guest_check_item_id/forward", restService.ForwardGuestCheckItem)
		guestCheckItem.Post("/:guest_check_item_id/deliver", restService.DeliverGuestCheckItem)
	}

	addr := fmt.Sprintf("0.0.0.0:%d", port)
	err := r.Listen(addr)
	if err != nil {
		log.Fatal("cannot start rest server", err)
	}

	log.Printf("rest server has been started on port %d", port)
}
