package routes

import (
	"sale-open-api/controllers"

	"github.com/labstack/echo/v4"
)

func RouteInit(e *echo.Group) {
	sale := e.Group("/sale")
	{
		sale.POST("", controllers.CreateSale)
		sale.GET("/:id", controllers.GetSaleByID)
		sale.GET("", controllers.GetSales)
		sale.DELETE("/:id", controllers.DeleteSaleByID)
		sale.PUT("/:id", controllers.UpdateSaleByID)
	}
	product := e.Group("/product")
	{
		product.GET("", controllers.GetProducts)
	}
}
