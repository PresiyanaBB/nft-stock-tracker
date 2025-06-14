package handlers

import (
	"context"
	"fmt"
	"time"

	"github.com/PresiyanaBB/nft-stock-tracker/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/skip2/go-qrcode"
)

type UserNFTHandler struct {
	repo models.UserNFTRepository
}

func (h *UserNFTHandler) GetManyUserNFTs(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	userId, _ := uuid.Parse(ctx.Locals("userId").(string))

	userNFTs, err := h.repo.GetManyUserNFTs(c, userId)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "User NFTs retrieved successfully",
		"data":    userNFTs,
	})
}

func (h *UserNFTHandler) GetAllUserNFTs(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	userNFTs, err := h.repo.GetAllUserNFTs(c)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "User NFTs retrieved successfully",
		"data":    userNFTs,
	})
}

func (h *UserNFTHandler) GetUserNFT(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	userNFTId, _ := uuid.Parse(ctx.Params("userNFT_id"))
	userId, _ := uuid.Parse(ctx.Locals("userId").(string))

	userNFT, err := h.repo.GetUserNFT(c, userId, userNFTId)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	var QRCode []byte
	QRCode, err = qrcode.Encode(
		fmt.Sprintf("userNFTId:%v,ownerId:%v", userNFTId, userId),
		qrcode.Medium,
		256,
	)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "User NFT  retrieved successfully",
		"data": &fiber.Map{
			"userNFT": userNFT,
			"qrcode":  QRCode,
		},
	})
}

func (h *UserNFTHandler) CreateUserNFT(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	userId, _ := uuid.Parse(ctx.Locals("userId").(string))
	userNFT := &models.UserNFT{}
	userNFT.UserID = userId

	if err := ctx.BodyParser(userNFT); err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
			"data":    nil,
		})
	}

	createdUserNFT, err := h.repo.CreateUserNFT(c, userId, userNFT)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
			"data":    nil,
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(&fiber.Map{
		"status":  "success",
		"message": "UserNFT created",
		"data":    createdUserNFT,
	})
}

func (h *UserNFTHandler) ValidateUserNFT(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	validateBody := &models.ValidateUserNFT{}

	if err := ctx.BodyParser(validateBody); err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
			"data":    nil,
		})
	}

	validateData := make(map[string]interface{})
	validateData["collected"] = true

	userNFT, err := h.repo.UpdateUserNFT(c, validateBody.OwnerId, validateBody.UserNFTId, validateData)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
			"data":    nil,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "Congratulations on your purchase",
		"data":    userNFT,
	})
}

func (h *UserNFTHandler) DeleteUserNFT(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	userNFTId, _ := uuid.Parse(ctx.Params("userNFT_id"))

	err := h.repo.DeleteUserNFT(c, userNFTId)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
			"data":    nil,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "User NFT deleted successfully",
		"data":    nil,
	})
}

func NewUserNFTHandler(router fiber.Router, repo models.UserNFTRepository) {
	handler := &UserNFTHandler{
		repo: repo,
	}

	router.Get("/", handler.GetManyUserNFTs)
	router.Get("/all", handler.GetAllUserNFTs)
	router.Post("/", handler.CreateUserNFT)
	router.Get("/:userNFT_id", handler.GetUserNFT)
	router.Post("/validate", handler.ValidateUserNFT)
	router.Delete("/:userNFT_id", handler.DeleteUserNFT)
}
