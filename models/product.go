package models

import "github.com/guregu/null"

type SaleOpenApiProduct struct {
	ID          int       `json:"id" gorm:"primary_key:auto_increment"`
	Name        string    `json:"name" gorm:"type: varchar(255)"`
	Description string    `json:"description" gorm:"type: text"`
	Price       float64   `json:"price" gorm:"type: float8"`
	CreatedAt   null.Time `json:"-"`
	UpdatedAt   null.Time `json:"-"`
}

func (SaleOpenApiProduct) TableName() string {
	return "sale_open_api_product"
}
