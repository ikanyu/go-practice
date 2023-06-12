package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	// "os"
)

var Database *gorm.DB

func Connect() {
	var err error
	// host := os.Getenv("DB_HOST")
	// username := os.Getenv("DB_USER")
	// password := os.Getenv("DB_PASSWORD")
	// databaseName := os.Getenv("DB_NAME")
	// port := os.Getenv("DB_PORT")

	dsn := "host=localhost user=wenyu password=xx dbname=diary_app port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	// dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Africa/Lagos", host, username, password, databaseName, port)
	Database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	fmt.Println(dsn)

	// pqDB, _ := Database.DB()
	// pqDB.Close()

	if err != nil {
			panic(err)
	} else {
			fmt.Println("Successfully connected to the database")

	}
}
