package routes

import (
	"net/http"

	"github.com/Dzikuri/shopifyx/internal/delivery/handler"
	"github.com/Dzikuri/shopifyx/internal/delivery/middleware"
	"github.com/labstack/echo/v4"
)

type RoutesConfig struct {
	Echo        *echo.Echo
	UserHandler *handler.UserHandler
    ProductHandler *handler.ProductHandler
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
	c.Echo.POST("/v1/user/login", c.UserHandler.UserLogin)
}

func (c *RoutesConfig) SetupRouteUser() {

}

func (c *RoutesConfig) SetupRouteProduct() {
    
	product := c.Echo.Group("/v1/product", func(next echo.HandlerFunc) echo.HandlerFunc {
        return middleware.JwtCheckTokenUser(next)
    })

	product.POST("", c.ProductHandler.ProductCreate)
}

func (c *RoutesConfig) SetupRouteBankAccount() {

}

func (c *RoutesConfig) SetupRoutePayment() {

}