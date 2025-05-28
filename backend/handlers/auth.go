package handlers

import (
	"context"
	// 	"fmt"
	"time"

	"github.com/PresiyanaBB/nft-stock-tracker/models"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

type AuthHandler struct {
	service models.AuthService
}

func (h *AuthHandler) Login(ctx *fiber.Ctx) error {
	creds := &models.AuthCredentials{}

	c, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := ctx.BodyParser(&creds); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "fail",
			"message": "Invalid request format. Please provide valid JSON.",
		})
	}

	if err := validate.Struct(creds); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "fail",
			"message": "Invalid credentials. Please provide a valid email and password.",
		})
	}

	token, user, err := h.service.Login(c, creds)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "fail",
			"message": "Invalid email or password.",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Successfully logged in",
		"data": fiber.Map{
			"token": token,
			"user":  user,
		},
	})
}

func (h *AuthHandler) Register(ctx *fiber.Ctx) error {
	creds := &models.AuthCredentials{}

	c, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := ctx.BodyParser(&creds); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "fail",
			"message": "Invalid request format. Please provide valid JSON.",
		})
	}

	if err := validate.Struct(creds); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "fail",
			"message": "Please provide a valid name, email, and password.",
		})
	}

	token, user, err := h.service.Register(c, creds)
	if err != nil {
		return ctx.Status(fiber.StatusConflict).JSON(fiber.Map{
			"status":  "fail",
			"message": "User already exists or registration failed.",
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  "success",
		"message": "Successfully registered",
		"data": fiber.Map{
			"token": token,
			"user":  user,
		},
	})
}

func NewAuthHandler(route fiber.Router, service models.AuthService) {
	handler := &AuthHandler{
		service: service,
	}

	route.Post("/login", handler.Login)
	route.Post("/register", handler.Register)
}
