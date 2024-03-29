package queries

import (
	"errors"

	"github.com/google/uuid"
	"github.com/oklave/collection-of-numbers/app/models"
	"gorm.io/gorm"
)

// Структура запросов gorm db
type NumberQueries struct {
	*gorm.DB
}

// Получить информацию по tg-id
func (q *NumberQueries) GetNumberById(id uuid.UUID) (models.Number, error) {
	// Определяем переменную
	number := models.Number{}

	// Send query to database.
	err := q.DB.Table("numbers", q.DB.Model(&Number)).Where("id = ?", id).Find(&Number).Error
	if err != nil {
		// вернуть объект с ошибкой
		return Number, errors.New("unable get user, DB error")
	}

	// Return query result.
	return number, nil
}
