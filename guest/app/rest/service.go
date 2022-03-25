package rest

import (
	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
	"github.com/patricksferraz/guest/domain/service"
)

type RestService struct {
	Service *service.Service
}

func NewRestService(service *service.Service) *RestService {
	return &RestService{
		Service: service,
	}
}

// CreateGuest godoc
// @Summary create a new guest
// @ID createGuest
// @Tags Guest
// @Description Router for create a new guest
// @Accept json
// @Produce json
// @Param body body CreateGuestRequest true "JSON body for create a new guest"
// @Success 200 {object} IDResponse
// @Failure 400 {object} HTTPResponse
// @Failure 403 {object} HTTPResponse
// @Router /guests [post]
func (t *RestService) CreateGuest(c *fiber.Ctx) error {
	var req CreateGuestRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(HTTPResponse{Msg: err.Error()})
	}

	guestID, err := t.Service.CreateGuest(c.Context(), &req.Name, &req.Doc)
	if err != nil {
		return c.Status(fiber.StatusForbidden).JSON(HTTPResponse{Msg: err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(IDResponse{ID: *guestID})
}

// FindGuest godoc
// @Summary find a gust
// @ID findGuest
// @Tags Guest
// @Description Router for find a guest
// @Accept json
// @Produce json
// @Param guest_id path string true "Guest ID"
// @Success 200 {object} Guest
// @Failure 400 {object} HTTPResponse
// @Failure 403 {object} HTTPResponse
// @Router /guests/{guest_id} [get]
func (t *RestService) FindGuest(c *fiber.Ctx) error {
	guestID := c.Params("guest_id")
	if !govalidator.IsUUIDv4(guestID) {
		return c.Status(fiber.StatusBadRequest).JSON(HTTPResponse{
			Msg: "guest_id is not a valid uuid",
		})
	}

	guest, err := t.Service.FindGuest(c.Context(), &guestID)
	if err != nil {
		return c.Status(fiber.StatusForbidden).JSON(HTTPResponse{Msg: err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(guest)
}
