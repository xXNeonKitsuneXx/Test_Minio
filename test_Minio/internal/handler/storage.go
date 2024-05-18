package handler

import (
	"test_Minio/internal/service"

	"github.com/gofiber/fiber/v2"
)

type StorageHandler struct {
	uploadSvc service.UploadService
}

func NewStorageHandler(uploadSvc service.UploadService) *StorageHandler {
	return &StorageHandler{
		uploadSvc: uploadSvc,
	}
}

func (h *StorageHandler) UploadFile(c *fiber.Ctx) error {
	file, err := c.FormFile("file")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "File not found",
		})
	}

	info, err := h.uploadSvc.UploadFile(file)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to upload file",
		})
	}
	return c.SendString(*info)
}
