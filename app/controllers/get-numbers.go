package controllers

import (
	"github.com/oklave/collection-of-numbers/pkg/response"
	"github.com/oklave/collection-of-numbers/platform/database"

	"github.com/oklave/fiber/v2"
)

func GetBooks(c *fiber.Ctx) error {
	// Вернуть все номера
	db := database.NumberDB()
	numbers, err := db.GetNumbers()
	if err != nil {
		// Если не найдено
		return response.RespondError(c, fiber.StatusNotFound, "numbers were not found")
	}

}
