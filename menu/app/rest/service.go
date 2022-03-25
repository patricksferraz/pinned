package rest

import (
	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
	"github.com/patricksferraz/menu/domain/service"
)

type RestService struct {
	Service *service.Service
}

func NewRestService(service *service.Service) *RestService {
	return &RestService{
		Service: service,
	}
}

// CreateMenu godoc
// @Summary create a new menu
// @ID createMenu
// @Tags Menu
// @Description Router for create a new menu
// @Accept json
// @Produce json
// @Param body body CreateMenuRequest true "JSON body for create a new menu"
// @Success 200 {object} IDResponse
// @Failure 400 {object} HTTPResponse
// @Failure 403 {object} HTTPResponse
// @Router /menus [post]
func (t *RestService) CreateMenu(c *fiber.Ctx) error {
	var req CreateMenuRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(HTTPResponse{Msg: err.Error()})
	}

	menuID, err := t.Service.CreateMenu(c.Context(), &req.Name)
	if err != nil {
		return c.Status(fiber.StatusForbidden).JSON(HTTPResponse{Msg: err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(IDResponse{ID: *menuID})
}

// FindMenu godoc
// @Summary find a menu
// @ID findMenu
// @Tags Menu
// @Description Router for find a menu
// @Accept json
// @Produce json
// @Param menu_id path string true "Menu ID"
// @Success 200 {object} Menu
// @Failure 400 {object} HTTPResponse
// @Failure 403 {object} HTTPResponse
// @Router /menus/{menu_id} [get]
func (t *RestService) FindMenu(c *fiber.Ctx) error {
	menuID := c.Params("menu_id")
	if !govalidator.IsUUIDv4(menuID) {
		return c.Status(fiber.StatusBadRequest).JSON(HTTPResponse{
			Msg: "menu_id is not a valid uuid",
		})
	}

	menu, err := t.Service.FindMenu(c.Context(), &menuID)
	if err != nil {
		return c.Status(fiber.StatusForbidden).JSON(HTTPResponse{Msg: err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(menu)
}

// CreateItem godoc
// @Summary create a new item
// @ID createItem
// @Tags Item
// @Description Router for create a new item
// @Accept json
// @Produce json
// @Param menu_id path string true "Menu ID"
// @Param body body CreateItemRequest true "JSON body for create a new item"
// @Success 200 {object} IDResponse
// @Failure 400 {object} HTTPResponse
// @Failure 403 {object} HTTPResponse
// @Router /menus/{menu_id}/items [post]
func (t *RestService) CreateItem(c *fiber.Ctx) error {
	var req CreateItemRequest

	menuID := c.Params("menu_id")
	if !govalidator.IsUUIDv4(menuID) {
		return c.Status(fiber.StatusBadRequest).JSON(HTTPResponse{
			Msg: "menu_id is not a valid uuid",
		})
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(HTTPResponse{Msg: err.Error()})
	}

	itemID, err := t.Service.CreateItem(c.Context(), &menuID, &req.Name, &req.Description, &req.Price, &req.Discount)
	if err != nil {
		return c.Status(fiber.StatusForbidden).JSON(HTTPResponse{Msg: err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(IDResponse{ID: *itemID})
}

// FindItem godoc
// @Summary find a item
// @ID findItem
// @Tags Item
// @Description Router for find a item
// @Accept json
// @Produce json
// @Param menu_id path string true "Menu ID"
// @Param item_id path string true "Item ID"
// @Success 200 {object} Item
// @Failure 400 {object} HTTPResponse
// @Failure 403 {object} HTTPResponse
// @Router /menus/{menu_id}/items/{item_id} [get]
func (t *RestService) FindItem(c *fiber.Ctx) error {
	menuID := c.Params("menu_id")
	if !govalidator.IsUUIDv4(menuID) {
		return c.Status(fiber.StatusBadRequest).JSON(HTTPResponse{
			Msg: "menu_id is not a valid uuid",
		})
	}

	itemID := c.Params("item_id")
	if !govalidator.IsUUIDv4(itemID) {
		return c.Status(fiber.StatusBadRequest).JSON(HTTPResponse{
			Msg: "item_id is not a valid uuid",
		})
	}

	item, err := t.Service.FindItem(c.Context(), &menuID, &itemID)
	if err != nil {
		return c.Status(fiber.StatusForbidden).JSON(HTTPResponse{Msg: err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(item)
}

// AddItemTag godoc
// @Summary add a item tag
// @ID addItemTag
// @Tags Item
// @Description Router for add a item tag
// @Accept json
// @Produce json
// @Param menu_id path string true "Menu ID"
// @Param item_id path string true "Item ID"
// @Param body body AddItemTagRequest true "JSON body for add a new item tag"
// @Success 200 {object} HTTPResponse
// @Failure 400 {object} HTTPResponse
// @Failure 403 {object} HTTPResponse
// @Router /menus/{menu_id}/items/{item_id}/tag [put]
func (t *RestService) AddItemTag(c *fiber.Ctx) error {
	var req AddItemTagRequest

	menuID := c.Params("menu_id")
	if !govalidator.IsUUIDv4(menuID) {
		return c.Status(fiber.StatusBadRequest).JSON(HTTPResponse{
			Msg: "menu_id is not a valid uuid",
		})
	}

	itemID := c.Params("item_id")
	if !govalidator.IsUUIDv4(itemID) {
		return c.Status(fiber.StatusBadRequest).JSON(HTTPResponse{
			Msg: "item_id is not a valid uuid",
		})
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(HTTPResponse{Msg: err.Error()})
	}

	err := t.Service.AddItemTags(c.Context(), &menuID, &itemID, &req.Tags)
	if err != nil {
		return c.Status(fiber.StatusForbidden).JSON(HTTPResponse{Msg: err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(HTTPResponse{Msg: "successful request"})
}
