package models

import "github.com/guregu/null"

type SaleOpenApiSaleHeader struct {
	ID              int       `json:"id" gorm:"primary_key:auto_increment"`
	TransactionDate null.Time `json:"transaction_date" gorm:"type: timestamptz"`
	InvoiceID       string    `json:"invoice_id" gorm:"type: varchar(50)"`
	CustomerName    string    `json:"customer_name" gorm:"type: varchar(255)"`
	Tax             string    `json:"tax" gorm:"type: varchar(50)"`
	Discount        string    `json:"discount" gorm:"type: varchar(50)"`
	Subtotal        float64   `json:"subtotal" gorm:"type: float8"`
	TotalPrice      float64   `json:"total_price" gorm:"type: float8"`
	CreatedAt       null.Time `json:"-"`
	UpdatedAt       null.Time `json:"-"`
}

func (SaleOpenApiSaleHeader) TableName() string {
	return "sale_open_api_sale_header"
}

type SaleOpenApiSaleDetail struct {
	ID        int       `json:"id" gorm:"primary_key:auto_increment"`
	ProductID int       `json:"product_id" gorm:"type: int8"`
	Quantity  int       `json:"quantity" gorm:"type: int8"`
	Price     float64   `json:"price" gorm:"type: float8"`
	HeaderID  int       `json:"header_id" gorm:"type: int8"`
	CreatedAt null.Time `json:"-"`
	UpdatedAt null.Time `json:"-"`
}

func (SaleOpenApiSaleDetail) TableName() string {
	return "sale_open_api_sale_detail"
}
