package utils

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

/* Запуск для prod

func StartServerWithGracefulShutdown(a *fiber.App) {

	Добавить канал для приема только shutdown! открыть его на всю функцию

	// Build Fiber connection URL.
	fiberConnURL, errConn := ConnectionURLBuilder("fiber")
	if errConn != nil {
		log.Fatal("fail to build fiber connection url")
	}

	// Run server.
	if err := a.Listen(fiberConnURL); err != nil {
		log.Printf("Oops... Server is not running! Reason: %v", err)
	}


}
*/

// Запуск для dev
func StartServer(a *fiber.App) {
	// Собираем обработку URL при помощи fiber
	fiberConnURL, _ := ConnectionURLBuilder("fiber")

	// Запускаем сервер
	if err := a.Listen(fiberConnURL); err != nil {
		log.Printf("Oops... Server is not running! Reason: %v", err)
	}
}
