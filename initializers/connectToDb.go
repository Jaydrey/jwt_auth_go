package initializers

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDb() {
	var err error
	db_user := os.Getenv("PG_USER")
	db_host := os.Getenv("PG_HOST")
	db_password := os.Getenv("PG_PASSWORD")
	db_name := os.Getenv("PG_NAME")
	db_port := os.Getenv("PG_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%v sslmode=disable ", db_host, db_user, db_password, db_name, db_port)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to DB")
	}

}
