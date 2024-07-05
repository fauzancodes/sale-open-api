package repositories

import (
	"sale-open-api/config"
	"sale-open-api/models"
)

func GetProducts(search string) (responses []models.SaleOpenApiProduct, err error) {
	var where string
	if search != "" {
		where += "name ILIKE '%" + search + "%'"
	}
	err = config.DB.Where(where).Find(&responses).Error

	return
}
