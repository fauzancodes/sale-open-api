package controllers

import (
	"net/http"
	"sale-open-api/repositories"
	"sale-open-api/res"
	"sale-open-api/utils"
	"strconv"
	"time"

	"github.com/guregu/null"
	"github.com/labstack/echo/v4"
)

func CreateSale(c echo.Context) error {
	var input res.SaleOpenApiSaleHeaderRequest
	if err := c.Bind(&input); err != nil {
		return c.JSON(400, utils.NewUnprocessableEntityError(err.Error()))
	}

	var transactionDate time.Time
	var err error
	if input.TransactionDate != "" {
		transactionDate, err = time.Parse("2006-01-02", input.TransactionDate)
		if err != nil {
			return c.JSON(500, utils.Respond(500, err, "Invalid transaction date format"))
		}
	}

	data, err := repositories.CreateSale(null.TimeFrom(transactionDate), &input)
	if err != nil {
		return c.JSON(500, utils.Respond(500, err, "Failed to create"))
	}

	for _, detailData := range input.Detail {
		repositories.CreateSaleDetail(data.ID, &detailData)
	}

	response, err := repositories.GetSaleByID(int(data.ID))
	if err != nil {
		return c.JSON(404, utils.Respond(404, err, "Failed to get response"))
	}

	return c.JSON(200, map[string]interface{}{
		"status":  200,
		"data":    response,
		"message": "Create Success",
	})
}

func GetSales(c echo.Context) error {
	param := utils.PopulatePaging(c, "")
	data := repositories.GetSales(param)

	return c.JSON(http.StatusOK, data)
}

func GetSaleByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := repositories.GetSaleByID(id)
	if err != nil {
		return c.JSON(404, utils.Respond(404, err, "Failed to get"))
	}
	return c.JSON(200, map[string]interface{}{
		"status":  200,
		"data":    data,
		"message": "Success to get",
	})
}

func DeleteSaleByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := repositories.GetSaleByIDPlain(id)
	if err != nil {
		return c.JSON(500, utils.Respond(500, err, "Failed to get"))
	}

	response, err := repositories.GetSaleByID(int(data.ID))
	if err != nil {
		return c.JSON(404, utils.Respond(404, err, "Failed to get response"))
	}

	detail, _ := repositories.GetAllSaleDetailByHeaderID(response.ID)
	for _, data := range detail {
		_, err = repositories.DeleteSaleDetail(data)
		if err != nil {
			return c.JSON(500, utils.Respond(500, err, "Failed to delete"))
		}
	}

	_, err = repositories.DeleteSale(data)
	if err != nil {
		return c.JSON(500, utils.Respond(500, err, "Failed to delete"))
	}

	return c.JSON(200, map[string]interface{}{
		"status":  200,
		"data":    response,
		"message": "Success to delete",
	})
}

func UpdateSaleByID(c echo.Context) error {
	var input res.SaleOpenApiSaleHeaderRequest
	if err := c.Bind(&input); err != nil {
		return c.JSON(400, utils.NewUnprocessableEntityError(err.Error()))
	}

	id, _ := strconv.Atoi(c.Param("id"))
	data, err := repositories.GetSaleByIDPlain(id)
	if err != nil {
		return c.JSON(500, utils.Respond(500, err, "Failed to get"))
	}

	if input.TransactionDate != "" {
		transactionDate, err := time.Parse("2006-01-02", input.TransactionDate)
		if err != nil {
			return c.JSON(500, utils.Respond(500, err, "Invalid transaction date format"))
		}
		data.TransactionDate = null.TimeFrom(transactionDate)
	}
	if input.InvoiceID != "" {
		data.InvoiceID = input.InvoiceID
	}
	if input.CustomerName != "" {
		data.CustomerName = input.CustomerName
	}
	if input.Tax != "" {
		data.Tax = input.Tax
	}
	if input.Discount != "" {
		data.Discount = input.Discount
	}
	if input.Subtotal > 0 {
		data.Subtotal = input.Subtotal
	}
	if input.TotalPrice > 0 {
		data.TotalPrice = input.TotalPrice
	}
	if len(input.Detail) > 0 {
		detail, _ := repositories.GetAllSaleDetailByHeaderID(data.ID)
		for _, detailData := range detail {
			repositories.DeleteSaleDetail(detailData)
		}
		for _, detailData := range input.Detail {
			repositories.CreateSaleDetail(data.ID, &detailData)
		}
	}

	dataUpdate, err := repositories.UpdateSale(data)
	if err != nil {
		return c.JSON(500, utils.Respond(500, err, "Failed to update"))
	}

	response, err := repositories.GetSaleByID(int(dataUpdate.ID))
	if err != nil {
		return c.JSON(404, utils.Respond(404, err, "Failed to get response"))
	}

	return c.JSON(200, map[string]interface{}{
		"status":  200,
		"data":    response,
		"message": "Success to update",
	})
}
