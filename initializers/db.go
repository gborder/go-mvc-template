package initializers

import (
	"fmt"
	"os"
	"v2/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func ConnectToDatabase() {
	var err error
	dsn := os.Getenv("DB_URL")
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Fail to connect to database")
	}
}

func SyncDB() {
	Db.AutoMigrate(&models.Post{})
}
