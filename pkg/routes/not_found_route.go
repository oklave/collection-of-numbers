package routes

import "github.com/gofiber/fiber/v2"

func NotFoundRoute(a *fiber.App) {
	// Register new special route.
	a.Use(
		// Anonymous function.
		func(c *fiber.Ctx) error {
			// Возвращает ошибку 404 и json ответ
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": true,
				"msg":   "endpoint is not found",
			})
		},
	)
}
