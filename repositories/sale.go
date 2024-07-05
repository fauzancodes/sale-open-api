package repositories

import (
	"sale-open-api/config"
	"sale-open-api/models"
	"sale-open-api/res"
	"sale-open-api/utils"

	"github.com/guregu/null"
)

func BuildSaleResponse(header models.SaleOpenApiSaleHeader) (response res.SaleOpenApiSaleHeaderResponse) {
	response.ID = header.ID
	response.TransactionDate = header.TransactionDate
	response.InvoiceID = header.InvoiceID
	response.CustomerName = header.CustomerName
	response.Tax = header.Tax
	response.Discount = header.Discount
	response.Subtotal = header.Subtotal
	response.TotalPrice = header.TotalPrice

	detail, _ := GetAllSaleDetailByHeaderID(response.ID)
	for _, data := range detail {
		response.Detail = append(response.Detail, res.SaleOpenApiSaleDetailResponse{
			ID:        data.ID,
			ProductID: data.ProductID,
			Quantity:  data.Quantity,
			Price:     data.Price,
		})
	}

	return
}

func CreateSale(transactionDate null.Time, data *res.SaleOpenApiSaleHeaderRequest) (response models.SaleOpenApiSaleHeader, err error) {
	response = models.SaleOpenApiSaleHeader{
		TransactionDate: transactionDate,
		InvoiceID:       data.InvoiceID,
		CustomerName:    data.CustomerName,
		Tax:             data.Tax,
		Discount:        data.Discount,
		Subtotal:        data.Subtotal,
		TotalPrice:      data.TotalPrice,
	}

	err = config.DB.Create(&response).Error
	if err != nil {
		return
	}

	return
}

func GetSales(param res.ReqPaging) (data res.ResPaging) {
	var responses []models.SaleOpenApiSaleHeader
	var where string
	if param.Search != "" {
		if where != "" {
			where += " AND "
		}
		where += "invoice_id ILIKE '%" + param.Search + "%'"
	}

	var modelTotal []models.SaleOpenApiSaleHeader

	var totalResult int64
	config.DB.Model(&modelTotal).Count(&totalResult)

	var totalFiltered int64
	config.DB.Model(&modelTotal).Where(where).Count(&totalFiltered)

	config.DB.Limit(param.Limit).Offset(param.Offset).Order(param.Sort + " " + param.Order).Where(where).Find(&responses)

	var responsesRefined []res.SaleOpenApiSaleHeaderResponse
	for _, data := range responses {
		responsesRefined = append(responsesRefined, BuildSaleResponse(data))
	}

	data = utils.PopulateResPaging(&param, responsesRefined, totalResult, totalFiltered)

	return
}

func GetSaleByID(id int) (response res.SaleOpenApiSaleHeaderResponse, err error) {
	var raw models.SaleOpenApiSaleHeader
	err = config.DB.First(&raw, id).Error

	response = BuildSaleResponse(raw)

	return
}

func GetSaleByIDPlain(id int) (response models.SaleOpenApiSaleHeader, err error) {
	err = config.DB.First(&response, id).Error

	return
}

func DeleteSale(request models.SaleOpenApiSaleHeader) (models.SaleOpenApiSaleHeader, error) {
	err := config.DB.Delete(&request).Error

	return request, err
}

func UpdateSale(request models.SaleOpenApiSaleHeader) (response models.SaleOpenApiSaleHeader, err error) {
	err = config.DB.Save(&request).Scan(&response).Error

	return
}

func CreateSaleDetail(headerID int, data *res.SaleOpenApiSaleDetailRequest) (response models.SaleOpenApiSaleDetail, err error) {
	response = models.SaleOpenApiSaleDetail{
		ProductID: data.ProductID,
		Quantity:  data.Quantity,
		Price:     data.Price,
		HeaderID:  headerID,
	}

	err = config.DB.Create(&response).Error
	if err != nil {
		return
	}

	return
}

func GetAllSaleDetailByHeaderID(headerID int) (response []models.SaleOpenApiSaleDetail, err error) {
	err = config.DB.Where("header_id = ?", headerID).Find(&response).Error

	return
}

func DeleteSaleDetail(request models.SaleOpenApiSaleDetail) (models.SaleOpenApiSaleDetail, error) {
	err := config.DB.Delete(&request).Error

	return request, err
}
