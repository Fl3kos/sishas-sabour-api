package database

import (
	"ShishaSabourAPI/internal/logger"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConectDB() *gorm.DB {
	//dsn := os.Getenv("DB_CONNECTION_STRING")

	db, err := gorm.Open(postgres.Open(""), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return db
}

func CloseDB(db *gorm.DB) {
	data, err := db.DB()
	if err != nil {
		logger.Fatalf("PANIC CLOSE THE DATABASE CONECTION")
	}
	data.Close()

}
