package routes

import (
	"net/http"

	"github.com/Dzikuri/shopifyx/internal/delivery/handler"
	"github.com/labstack/echo/v4"
)

type RoutesConfig struct {
	Echo        *echo.Echo
	UserHandler *handler.UserHandler
}

func (c *RoutesConfig) Setup() {

	c.Echo.GET("/health", func(c echo.Context) error {
		// Your existing health-check handler logic
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Ok",
		})
	})

	// Error Handling Page not found
	c.Echo.Any("/*", func(c echo.Context) error {
		var d struct{}
		return c.JSON(http.StatusNotFound, map[string]interface{}{"message": "Not found", "data": d})
	})

	c.SetupRouteAuth()
	c.SetupRouteUser()
	c.SetupRouteProduct()
	c.SetupRouteBankAccount()
	c.SetupRoutePayment()
}

func (c *RoutesConfig) SetupRouteAuth() {
	c.Echo.POST("/v1/user/register", c.UserHandler.UserRegister)
}

func (c *RoutesConfig) SetupRouteUser() {

}

func (c *RoutesConfig) SetupRouteProduct() {

}

func (c *RoutesConfig) SetupRouteBankAccount() {

}

func (c *RoutesConfig) SetupRoutePayment() {

}
