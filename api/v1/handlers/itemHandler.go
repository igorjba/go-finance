package handlers

import (
	"context"

	"github.com/igorjba/go-test-back/repository/item"
	"github.com/igorjba/go-test-back/service"

	"github.com/gofiber/fiber/v2"
)

type ItemHandler struct {
	Service *service.ItemService
}

func NewItemHandler(s *service.ItemService) *ItemHandler {
	return &ItemHandler{Service: s}
}

func (h *ItemHandler) GetAllItems(c *fiber.Ctx) error {
	items, err := h.Service.GetAll(context.TODO())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(newErrorResponse("Failed to retrieve items"))
	}
	return c.Status(fiber.StatusOK).JSON(items)
}

func (h *ItemHandler) GetItemByID(c *fiber.Ctx) error {
	id := c.Params("id")
	item, err := h.Service.GetByID(context.TODO(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(newErrorResponse("Failed to retrieve item"))
	}
	return c.Status(fiber.StatusOK).JSON(item)
}

func (h *ItemHandler) CreateItem(c *fiber.Ctx) error {
	var item item.Item
	if err := c.BodyParser(&item); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(newErrorResponse("Failed to parse request"))
	}

	err := h.Service.Create(context.TODO(), &item)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(newErrorResponse("Failed to create item"))
	}

	return c.Status(fiber.StatusCreated).JSON(item)
}

func (h *ItemHandler) UpdateItem(c *fiber.Ctx) error {
	id := c.Params("id")
	var updatedItem item.Item
	if err := c.BodyParser(&updatedItem); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(newErrorResponse("Failed to parse request"))
	}

	updatedItem.ID = id
	err := h.Service.Update(context.TODO(), id, &updatedItem)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(newErrorResponse("Failed to update item"))
	}

	return c.Status(fiber.StatusOK).JSON(updatedItem)
}

func (h *ItemHandler) DeleteItem(c *fiber.Ctx) error {
	id := c.Params("id")
	err := h.Service.Delete(context.TODO(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(newErrorResponse("Failed to delete item"))
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Item deleted successfully"})
}
