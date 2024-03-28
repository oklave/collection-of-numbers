package controllers

import (
	"github.com/oklave/collection-of-numbers/pkg/response"

	"github.com/oklave/fiber/v2"
)

func GetNumber(c *fiber.Ctx) error {
	// Выдергиваем tg-id ID from URL.
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return response.RespondError(c, fiber.StatusInternalServerError, err.Error())
	}

}
