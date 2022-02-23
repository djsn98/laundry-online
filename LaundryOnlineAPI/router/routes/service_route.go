package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var ServiceRoute = []Route{
	{
		URI:     "/service",
		Method:  http.MethodPost,
		Handler: func(c *gin.Context) { serviceController.Create(c) },
	},
	{
		URI:     "/services",
		Method:  http.MethodGet,
		Handler: func(c *gin.Context) { serviceController.ReadAll(c) },
	},
	{
		URI:     "/service/:service_id",
		Method:  http.MethodGet,
		Handler: func(c *gin.Context) { serviceController.ReadById(c) },
	},
	{
		URI:     "/service/:service_id",
		Method:  http.MethodPut,
		Handler: func(c *gin.Context) { serviceController.Update(c) },
	},
	{
		URI:     "/service/:service_id",
		Method:  http.MethodDelete,
		Handler: func(c *gin.Context) { serviceController.Delete(c) },
	},
}
