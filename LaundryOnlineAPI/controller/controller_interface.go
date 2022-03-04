package controller

import "github.com/gin-gonic/gin"

type ServiceControllerInterface interface {
	Create(c *gin.Context)
	ReadAll(c *gin.Context)
	ReadById(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type CustomerControllerInterface interface {
	Create(c *gin.Context)
	ReadAll(c *gin.Context)
	ReadByUsername(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type OrderControllerInterface interface {
	Create(c *gin.Context)
	ReadAll(c *gin.Context)
	ReadById(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type LoginCustomerControllerInterface interface {
	Login(c *gin.Context)
}
