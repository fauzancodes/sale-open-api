package controllers

import (
	"sale-open-api/repositories"
	"sale-open-api/utils"

	"github.com/labstack/echo/v4"
)

func GetProducts(c echo.Context) error {
	search := c.QueryParam("search")
	data, err := repositories.GetProducts(search)
	if err != nil {
		return c.JSON(404, utils.Respond(404, err, "Failed to get"))
	}

	return c.JSON(200, map[string]interface{}{
		"status":  200,
		"data":    data,
		"message": "Success to get",
	})
}
