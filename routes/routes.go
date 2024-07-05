package routes

import (
	"net/http"
	"sale-open-api/controllers"

	"github.com/labstack/echo/v4"
)

func RouteInit(e *echo.Group) {
	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "Welcome to nizom sale open API")
	})
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
