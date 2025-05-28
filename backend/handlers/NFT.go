package handlers

import (
	"context"
	"time"

	"github.com/PresiyanaBB/nft-stock-tracker/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type NFTHandler struct {
	repo models.NFTRepository
}

func (h *NFTHandler) GetManyNFTs(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()
	nfts, err := h.repo.GetManyNFTs(c)

	if err != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "NFTs retrieved successfully",
		"data":    nfts,
	})
}

func (h *NFTHandler) GetNFT(ctx *fiber.Ctx) error {
	nftId, _ := uuid.Parse(ctx.Params("id"))

	c, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	nft, err := h.repo.GetNFT(c, nftId)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "NFT retrieved successfully",
		"data":    nft,
	})
}

func (h *NFTHandler) CreateNFT(ctx *fiber.Ctx) error {
	nft := &models.NFT{}

	c, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	if err := ctx.BodyParser(nft); err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
			"data":    nil,
		})
	}

	createdNFT, err := h.repo.CreateNFT(c, nft)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
			"data":    nil,
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(&fiber.Map{
		"status":  "success",
		"message": "NFT created successfully",
		"data":    createdNFT,
	})
}

func (h *NFTHandler) UpdateNFT(ctx *fiber.Ctx) error {
	nftId, _ := uuid.Parse(ctx.Params("id"))
	updatedNFT := make(map[string]interface{})

	c, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	if err := ctx.BodyParser(&updatedNFT); err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	nft, err := h.repo.UpdateNFT(c, nftId, updatedNFT)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
			"data":    nil,
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(&fiber.Map{
		"status":  "success",
		"message": "NFT created successfully",
		"data":    nft,
	})

}

func (h *NFTHandler) DeleteNFT(ctx *fiber.Ctx) error {
	nftId, _ := uuid.Parse(ctx.Params("id"))

	c, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	err := h.repo.DeleteNFT(c, nftId)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return ctx.SendStatus(fiber.StatusNoContent)
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
