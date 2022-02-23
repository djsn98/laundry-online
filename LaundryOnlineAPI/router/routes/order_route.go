package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var OrderRoute = []Route{
	{
		URI:     "/order",
		Method:  http.MethodPost,
		Handler: func(c *gin.Context) { orderController.Create(c) },
	},
	{
		URI:     "/orders",
		Method:  http.MethodGet,
		Handler: func(c *gin.Context) { orderController.ReadAll(c) },
	},
	{
		URI:     "/order/:order_id",
		Method:  http.MethodGet,
		Handler: func(c *gin.Context) { orderController.ReadById(c) },
	},
	{
		URI:     "/order/:order_id",
		Method:  http.MethodPut,
		Handler: func(c *gin.Context) { orderController.Update(c) },
	},
	{
		URI:     "/order/:order_id",
		Method:  http.MethodDelete,
		Handler: func(c *gin.Context) { orderController.Delete(c) },
	},
}
