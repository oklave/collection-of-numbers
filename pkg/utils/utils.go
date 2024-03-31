package utils

import (
	"fmt"
	"os"
)

// создает url-соединение
func ConnectionURLBuilder(str string) (string, error) {
	// определяем переменную для url-соединения
	var url string

	// url-соединение
	switch str {
	case "postgres":
		url = fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=%s",
			os.Getenv("POSTGRES_HOST"),
			os.Getenv("POSTGRES_PORT"),
			os.Getenv("POSTGRES_USER"),
			os.Getenv("POSTGRES_PASSWORD"),
			os.Getenv("POSTGRES_NAME"),
			os.Getenv("POSTGRES_SSL_MODE"),
			os.Getenv("POSTGRES_TIME_ZONE"),
		)
	case "redis":
		url = fmt.Sprintf(
			"%s:%s",
			os.Getenv("REDIS_HOST"),
			os.Getenv("REDIS_PORT"),
		)
	case "fiber":
		url = fmt.Sprintf(
			"%s:%s",
			os.Getenv("SERVER_HOST"),
			os.Getenv("SERVER_PORT"),
		)
	default:
		return "", fmt.Errorf("connection name '%v' is not supported", str)
	}

	return url, nil
}
