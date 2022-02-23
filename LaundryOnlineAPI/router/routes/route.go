package routes

import "github.com/gin-gonic/gin"

type route struct {
	URI     string
	Method  string
	Handler func(c *gin.Context)
}
