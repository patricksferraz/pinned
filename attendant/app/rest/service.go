package rest

import (
	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
	"github.com/patricksferraz/attendant/domain/service"
)

type RestService struct {
	Service *service.Service
}

func NewRestService(service *service.Service) *RestService {
	return &RestService{
		Service: service,
	}
}

// CreateAttendant godoc
// @Summary create a new attendant
// @ID createAttendant
// @Tags Attendant
// @Description Router for create a new attendant
// @Accept json
// @Produce json
// @Param body body CreateAttendantRequest true "JSON body for create a new attendant"
// @Success 200 {object} IDResponse
// @Failure 400 {object} HTTPResponse
// @Failure 403 {object} HTTPResponse
// @Router /attendants [post]
func (t *RestService) CreateAttendant(c *fiber.Ctx) error {
	var req CreateAttendantRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(HTTPResponse{Msg: err.Error()})
	}

	attendantID, err := t.Service.CreateAttendant(c.Context(), &req.Name)
	if err != nil {
		return c.Status(fiber.StatusForbidden).JSON(HTTPResponse{Msg: err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(IDResponse{ID: *attendantID})
}

// FindAttendant godoc
// @Summary find a attendant
// @ID findAttendant
// @Tags Attendant
// @Description Router for find a attendant
// @Accept json
// @Produce json
// @Param attendant_id path string true "Attendant ID"
// @Success 200 {object} Attendant
// @Failure 400 {object} HTTPResponse
// @Failure 403 {object} HTTPResponse
// @Router /attendants/{attendant_id} [get]
func (t *RestService) FindAttendant(c *fiber.Ctx) error {
	attendantID := c.Params("attendant_id")
	if !govalidator.IsUUIDv4(attendantID) {
		return c.Status(fiber.StatusBadRequest).JSON(HTTPResponse{
			Msg: "attendant_id is not a valid uuid",
		})
	}

	attendant, err := t.Service.FindAttendant(c.Context(), &attendantID)
	if err != nil {
		return c.Status(fiber.StatusForbidden).JSON(HTTPResponse{Msg: err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(attendant)
}
