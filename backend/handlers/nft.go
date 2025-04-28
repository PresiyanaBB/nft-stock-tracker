package handlers

import (
	"context"
	"github.com/PresiyanaBB/crypto-price-tracker/models"
	"github.com/gofiber/fiber/v2"
	"time"
)

type NFTHandler struct {
	repo models.NFTRepository
}

func (h *NFTHandler) GetManyNFTs(ctx *fiber.Ctx) error {
	context, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()
	nfts, err := h.repo.GetManyNFTs(context)

	if err != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "success",
		"data":   nfts,
	})
}

func (h *NFTHandler) GetNFT(c *fiber.Ctx) error {
	return nil
}

func (h *NFTHandler) CreateNFT(c *fiber.Ctx) error {
	return nil
}

func (h *NFTHandler) AddNFT(c *fiber.Ctx) error {
	return nil
}

func (h *NFTHandler) UpdateNFT(c *fiber.Ctx) error {
	return nil
}

func (h *NFTHandler) DeleteNFT(c *fiber.Ctx) error {
	return nil
}

func NewNFTHandler(router fiber.Router, repo models.NFTRepository) {
	handler := &NFTHandler{
		repo: repo,
	}

	router.Get("/", handler.GetManyNFTs)
	router.Get("/:id", handler.GetNFT)
	router.Post("/", handler.CreateNFT)
	router.Put("/:id", handler.UpdateNFT)
	router.Delete("/:id", handler.DeleteNFT)
}
