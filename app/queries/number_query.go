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
	// Определяем модель
	number := models.Number{}

	// Отправляем запрос к таблице на такой-то id
	err := q.DB.Table("numbers", q.DB.Model(&number)).Where("id = ?", id).Find(&number).Error
	if err != nil {
		// Вернуть объект с ошибкой
		return number, errors.New("unable get user, DB error")
	}

	// Вернуть данные по tg-id
	return number, nil
}
