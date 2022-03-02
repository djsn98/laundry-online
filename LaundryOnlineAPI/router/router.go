package router

import (
	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	// Instance new gin router and setup router
	router := gin.Default()
	return setupRouter(router)
}
