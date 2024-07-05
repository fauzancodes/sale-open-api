package handler

import (
	"fmt"
	"net/http"
	"os"
	"sale-open-api/config"
	"sale-open-api/routes"

	_ "github.com/joho/godotenv/autoload"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := Start()

	var port = os.Getenv("PORT")

	fmt.Println("server running localhost:" + port)
	e.Logger.Fatal(e.Start(":" + port))
}

func Main(w http.ResponseWriter, r *http.Request) {
	e := Start()

	e.ServeHTTP(w, r)
}

func Start() *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
		AllowHeaders: []string{"X-Requested-With", "Content-Type", "Authorization"},
	}))

	config.DatabaseInit()
	config.RunMigration()

	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "Welcome to nizom sale open API")
	})
	routes.RouteInit(e.Group("/api"))

	return e
}
