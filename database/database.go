package database

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/rest-go/database/models"

	// Import the postgres dialect.
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func InitDb() *gorm.DB {

	db := openConnection()

	db.LogMode(true)
	models.RunMigrations(db)

	return db
}

func openConnection() *gorm.DB {
	maxRetries := 3
	waitTime := 5 // In seconds

	for i := 1; i <= maxRetries; i++ {
		fmt.Printf("Opening Connection; Attempt %d of %d...\n", i, maxRetries)
		dbConfig := "sslmode=disable host=db port=5432 dbname=todos user=tduser password=tdpass"

		db, err := gorm.Open("postgres", dbConfig)
		if err != nil {
			fmt.Printf("Cannot open connection (retrying in %ds): %v\n", waitTime, err)
			time.Sleep(time.Duration(waitTime) * time.Second)
			continue
		}

		return db
	}

	panic(fmt.Errorf("Cannot open database connection after %d retries!\n", maxRetries))
}
