package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var CustomerRoute = []Route{
	{
		URI:     "/customer",
		Method:  http.MethodPost,
		Handler: func(c *gin.Context) { customerController.Create(c) },
	},
	{
		URI:     "/customers",
		Method:  http.MethodGet,
		Handler: func(c *gin.Context) { customerController.ReadAll(c) },
	},
	{
		URI:     "/customer/:customer_username",
		Method:  http.MethodGet,
		Handler: func(c *gin.Context) { customerController.ReadByUsername(c) },
	},
	{
		URI:     "/customer/:customer_username",
		Method:  http.MethodPut,
		Handler: func(c *gin.Context) { customerController.Update(c) },
	},
	{
		URI:     "/customer/:customer_username",
		Method:  http.MethodDelete,
		Handler: func(c *gin.Context) { customerController.Delete(c) },
	},
}
