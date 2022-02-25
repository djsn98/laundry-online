package router

import (
	"LaundryOnlineAPI/middlewares"
	"LaundryOnlineAPI/router/routes"
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupRouter(router *gin.Engine) *gin.Engine {
	router.Use(middlewares.DisableCORSMiddleware())
	routes := routes.Load()

	for _, route := range routes {
		switch route.Method {
		case http.MethodPost:
			router.POST(route.URI, route.Handler)
		case http.MethodGet:
			router.GET(route.URI, route.Handler)
		case http.MethodPut:
			router.PUT(route.URI, route.Handler)
		case http.MethodDelete:
			router.DELETE(route.URI, route.Handler)
		default:
			panic("Http method not available")
		}
	}

	return router
}
