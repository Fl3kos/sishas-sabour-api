package service

import (
	"ShishaSabourAPI/internal/database"
	"ShishaSabourAPI/internal/models"
)

// TODO: method to get info of database
func GetTobbacosDB() []models.Tobacco {
	var tobaccos []models.Tobacco
	db := database.ConectDB()

	defer database.CloseDB(db)

	db.Find(&tobaccos)

	return tobaccos
}

// TODO: Method to post info on database
