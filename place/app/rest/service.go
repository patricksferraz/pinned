package rest

import (
	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
	"github.com/patricksferraz/place/domain/service"
)

type RestService struct {
	Service *service.Service
}

func NewRestService(service *service.Service) *RestService {
	return &RestService{
		Service: service,
	}
}

// CreatePlace godoc
// @Summary create a new place
// @ID createPlace
// @Tags Place
// @Description Router for create a new place
// @Accept json
// @Produce json
// @Param body body CreatePlaceRequest true "JSON body for create a new place"
// @Success 200 {object} IDResponse
// @Failure 400 {object} HTTPResponse
// @Failure 403 {object} HTTPResponse
// @Router /places [post]
func (t *RestService) CreatePlace(c *fiber.Ctx) error {
	var req CreatePlaceRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(HTTPResponse{Msg: err.Error()})
	}

	placeID, err := t.Service.CreatePlace(c.Context(), &req.Name)
	if err != nil {
		return c.Status(fiber.StatusForbidden).JSON(HTTPResponse{Msg: err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(IDResponse{ID: *placeID})
}

// FindPlace godoc
// @Summary find a place
// @ID findPlace
// @Tags Place
// @Description Router for find a place
// @Accept json
// @Produce json
// @Param place_id path string true "Place ID"
// @Success 200 {object} Place
// @Failure 400 {object} HTTPResponse
// @Failure 403 {object} HTTPResponse
// @Router /places/{place_id} [get]
func (t *RestService) FindPlace(c *fiber.Ctx) error {
	placeID := c.Params("place_id")
	if !govalidator.IsUUIDv4(placeID) {
		return c.Status(fiber.StatusBadRequest).JSON(HTTPResponse{
			Msg: "place_id is not a valid uuid",
		})
	}

	place, err := t.Service.FindPlace(c.Context(), &placeID)
	if err != nil {
		return c.Status(fiber.StatusForbidden).JSON(HTTPResponse{Msg: err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(place)
}
