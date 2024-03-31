package database

import (
	"fmt"
	"log"

	"github.com/oklave/collection-of-numbers/app/queries"
	"github.com/oklave/collection-of-numbers/pkg/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

// Queries struct for collect all app queries.
type Queries struct {
	*queries.UserQueries // load queries from User model
	*queries.BookQueries // load queries from Book model
}

// InitDBConnection func for connection to PostgreSQL database.
func InitDBConnection() (*Queries, error) {
	// postgresql url-соединение
	postgresConnURL, errCreateDBUrl := utils.ConnectionURLBuilder("postgres")
	if errCreateDBUrl != nil {
		return nil, errCreateDBUrl
	}

	// Define database connection for PostgreSQL.
	db, err = gorm.Open(postgres.Open(postgresConnURL), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error, not connected to database, %w", err)
	}

	return &Queries{
		NumberQueries: &queries.NumberQueries{DB: db},
	}, nil
}

func CloseDBConnection(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		log.Println("fail catch db connection")
	}
	err = sqlDB.Close()
	if err != nil {
		log.Fatal("fail close db connection")
	}
	return nil
}

// NumberDB used for init number db query
func NumberDB() *queries.NumberQueries {
	return &queries.NumberQueries{DB: db}
}
