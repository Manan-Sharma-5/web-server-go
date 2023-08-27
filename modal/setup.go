package modal

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=localhost user=postgres dbname=go_blog port=5432 sslmode=disable timezone=Asia/Kolkata"
	database, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		panic("Failed to connect to Database!")
	}
	database.AutoMigrate(&Post{})
	DB = database
}
