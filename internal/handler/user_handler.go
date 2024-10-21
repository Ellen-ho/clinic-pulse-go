package handler

import (
    "clinic-pulse-go/internal/domain"
    "clinic-pulse-go/internal/usecase"
    "github.com/gofiber/fiber/v2"
)

type UserHandler struct {
    usecase usecase.UserUsecase
}

func NewUserHandler(usecase usecase.UserUsecase) *UserHandler {
    return &UserHandler{usecase: usecase}
}

func (h *UserHandler) GetUser(c *fiber.Ctx) error {
    id, _ := c.ParamsInt("id")
    user, err := h.usecase.GetUser(id)
    if err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
    }
    return c.JSON(user)
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
    user := new(domain.User)
    if err := c.BodyParser(user); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
    }
    err := h.usecase.CreateUser(user)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Cannot create user"})
    }
    return c.Status(fiber.StatusCreated).JSON(user)
}