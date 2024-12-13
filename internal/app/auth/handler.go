package auth

import (
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

type SignUpRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type SignInRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type RecoveryRequest struct {
	Email string `json:"email" validate:"required"`
}

type ResetPasswordRequest struct {
	Password   string `json:"password" validate:"required"`
	RePassword string `json:"re_password" validate:"required"`
}

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

var validate = validator.New()

func (h *Handler) SignUp(c *fiber.Ctx) error {
	var req SignUpRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Invalid request body",
			"error":   err.Error(),
		})
	}

	if err := validate.Struct(req); err != nil {
		validationErrors := make(map[string]string)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Validation failed",
			"error":   validationErrors,
		})
	}

	resp, err := h.service.SignUpService(c.Context(), req.Name, req.Email, req.Password)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "error.code",
			"error":   err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":   true,
		"message":  "http.ok",
		"response": c.JSON(resp),
	})
}

func (h *Handler) SignIn(c *fiber.Ctx) error {
	var req SignInRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Invalid request body",
			"error":   err.Error(),
		})
	}

	if err := validate.Struct(req); err != nil {
		validationErrors := make(map[string]string)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Validation failed",
			"error":   validationErrors,
		})
	}

	resp, err := h.service.SignInService(c.Context(), req.Email, req.Password)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "error.code",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":   true,
		"message":  "http.ok",
		"response": c.JSON(resp),
	})
}

func (h *Handler) Recovery(c *fiber.Ctx) error {
	var req RecoveryRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Invalid request body",
			"error":   err.Error(),
		})
	}

	if err := validate.Struct(req); err != nil {
		validationErrors := make(map[string]string)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Validation failed",
			"error":   validationErrors,
		})
	}

	resp, err := h.service.RecoveryService(c.Context(), req.Email)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":   true,
		"message":  "http.ok",
		"response": c.JSON(resp),
	})
}

func (h *Handler) ResetPassword(c *fiber.Ctx) error {
	var req ResetPasswordRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Invalid request body",
			"error":   err.Error(),
		})
	}

	if err := validate.Struct(req); err != nil {
		validationErrors := make(map[string]string)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Validation failed",
			"error":   validationErrors,
		})
	}

	resp, err := h.service.ResetPasswordService(c.Context(), c.Params("id"), req.Password)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"status":   true,
		"message":  "http.ok",
		"response": c.JSON(resp),
	})
}
