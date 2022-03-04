package controller

import (
	"LaundryOnlineAPI/model/web"
	"LaundryOnlineAPI/model/web/customerReqRes"
	"LaundryOnlineAPI/usecase"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginCustomerControllerImpl struct {
	LoginCustomerUsecase usecase.CustomerLoginUsecaseInterface
}

func NewLoginCustomerControllerImpl(LoginCustomerUsecase usecase.CustomerLoginUsecaseInterface) LoginCustomerControllerInterface {
	return &LoginCustomerControllerImpl{LoginCustomerUsecase: LoginCustomerUsecase}
}

func (lcci *LoginCustomerControllerImpl) Login(c *gin.Context) {
	var loginCustomerReq customerReqRes.LoginCustomerReq

	err := c.ShouldBindJSON(&loginCustomerReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   web.MessageRes{Message: err.Error()},
		})
		return
	}
	tokenString, ok, err := lcci.LoginCustomerUsecase.Login(context.Background(), &loginCustomerReq)
	if (err != nil) || !ok {
		c.JSON(http.StatusInternalServerError, web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL SERVER ERROR",
			Data:   web.MessageRes{Message: err.Error()},
		})
		return
	}

	c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   web.MessageRes{Message: "Token " + tokenString},
	})
	return
}
