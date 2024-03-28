package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/oklave/collection-of-numbers/app/controllers"
	"github.com/oklave/collection-of-numbers/pkg/middleware"
)

// Функция маршрутизации url запросов начинающихся с /number
func NumberRoutes(a *fiber.App) {
	// Создаем группу
	route := a.Group("/number")

	// Маршруты для POST метода:
	// route.Post("/tg-id", middleware.JWTProtected(), controllers.CreateNumber) // create a new book

	// Маршруты для PUT метода:
	// route.Put("/tg-id/:id", middleware.JWTProtected(), controllers.UpdateNumber) // update one book by ID

	// Маршруты для DELETE метода:
	// route.Delete("/tg-id/:id", middleware.JWTProtected(), controllers.DeleteNumber) // delete one book by ID

	// Маршруты GET метода:
	route.Get("/all-tg-id", middleware.BasicAuth(), controllers.GetNumber) // получить все тг-id
	route.Get("/tg-id/:id", middleware.BasicAuth(), controllers.GetNumber) // получить определенный тг-id
}
