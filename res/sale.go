package res

import "github.com/guregu/null"

type SaleOpenApiSaleHeaderRequest struct {
	TransactionDate string                         `json:"transaction_date"`
	InvoiceID       string                         `json:"invoice_id"`
	CustomerName    string                         `json:"customer_name"`
	Tax             string                         `json:"tax"`
	Discount        string                         `json:"discount"`
	Subtotal        float64                        `json:"subtotal"`
	TotalPrice      float64                        `json:"total_price"`
	Detail          []SaleOpenApiSaleDetailRequest `json:"detail"`
}

type SaleOpenApiSaleDetailRequest struct {
	ProductID int     `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
}

type SaleOpenApiSaleHeaderResponse struct {
	ID              int                             `json:"id"`
	TransactionDate null.Time                       `json:"transaction_date"`
	InvoiceID       string                          `json:"invoice_id"`
	CustomerName    string                          `json:"customer_name"`
	Tax             string                          `json:"tax"`
	Discount        string                          `json:"discount"`
	Subtotal        float64                         `json:"subtotal"`
	TotalPrice      float64                         `json:"total_price"`
	Detail          []SaleOpenApiSaleDetailResponse `json:"detail"`
}

type SaleOpenApiSaleDetailResponse struct {
	ID        int     `json:"id"`
	ProductID int     `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
}
