package main

import (
	"github.com/oklave/collection-of-numbers/pkg/configs"
	//"github.com/oklave/collection-of-numbers/pkg/middleware"
	"github.com/oklave/collection-of-numbers/pkg/routes"
	"github.com/oklave/collection-of-numbers/pkg/utils"
	"github.com/oklave/collection-of-numbers/platform/database"

	//"github.com/oklave/collection-of-numbers/platform/migrations"
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Подгружаем конфиг fiber
	// Из переменной среды берем
	// 1) SERVER_READ_TIMEOUT
	config := configs.FiberConfig()

	// Определяем новое приложение fiber с нужным конфигом
	app := fiber.New(config)

	// middlewares
	//middleware.FiberMiddleware(app)

	// проверка наличия файла с переменными среды
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	// инициализация соединения с БД
	_, errInitDb := database.InitDBConnection()
	if errInitDb != nil {
		log.Fatal("database not load")
	}
	/*
		// миграции .. to soon
		migrationFileSource := os.Getenv("SQL_SOURCE_PATH")
		err = migrations.Migrate(migrationFileSource)
		if err != nil {
			log.Fatal("database migration fail")
		}
	*/
	// маршрутизация url
	//routes.SwaggerRoute(app)   // Документация api (swagger)
	// routes.UsersRoutes(app)
	routes.NumberRoutes(app)  // Поиск номеров api
	routes.NotFoundRoute(app) // 404 Error.

	// запуск сервера в dev/prod режиме
	if os.Getenv("STAGE_STATUS") == "dev" {
		utils.StartServer(app)
		//} else {
		//	utils.StartServerWithGracefulShutdown(app)
	}
}
