package controllers

import (
	"github.com/google/uuid"
	"github.com/oklave/collection-of-numbers/pkg/response"
	"github.com/oklave/collection-of-numbers/platform/database"

	"github.com/gofiber/fiber/v2"
)

func GetNumber(c *fiber.Ctx) error {
	// Выдергиваем tg-id ID from URL.
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return response.RespondError(c, fiber.StatusInternalServerError, err.Error())
	}
	db := database.NumberDB()
	book, err := db.GetBookById(id)
	if err != nil {
		// Return, if book not found.
		return response.RespondError(c, fiber.StatusNotFound, "book with the given ID is not found")
	}

	// Return status 200 OK.
	return response.RespondSuccess(c, fiber.StatusOK, book)
}
