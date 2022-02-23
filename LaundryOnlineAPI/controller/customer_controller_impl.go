package controller

import (
	"LaundryOnlineAPI/model/web"
	"LaundryOnlineAPI/model/web/customerReqRes"
	"LaundryOnlineAPI/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CustomerControllerImpl struct {
	CustomerUsecase usecase.CustomerCRUDUsecaseInterface
}

func NewCustomerControllerImpl(CustomerUsecase usecase.CustomerCRUDUsecaseInterface) CustomerControllerInterface {
	return &CustomerControllerImpl{CustomerUsecase: CustomerUsecase}
}

func (cci *CustomerControllerImpl) Create(c *gin.Context) {
	var createCustomerReq customerReqRes.CreateCustomerReq

	err := c.ShouldBindJSON(&createCustomerReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   web.MessageRes{Message: err.Error()},
		})
		return
	}

	err2 := cci.CustomerUsecase.Create(c.Request.Context(), &createCustomerReq)
	if err2 != nil {
		c.JSON(http.StatusInternalServerError, web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL SERVER ERROR",
			Data:   web.MessageRes{Message: err2.Error()},
		})
		return
	}

	c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   web.MessageRes{Message: "Success add new customer!"},
	})
	return
}

func (cci *CustomerControllerImpl) ReadAll(c *gin.Context) {
	result, err := cci.CustomerUsecase.ReadAll(c.Request.Context())
	if err.Error() != "Customer not found!" {
		c.JSON(http.StatusNotFound, web.WebResponse{
			Code:   http.StatusNotFound,
			Status: "NOT FOUND",
			Data:   web.MessageRes{Message: err.Error()},
		})
		return
	}
	if err != nil {
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
		Data:   *result,
	})
	return
}

func (cci *CustomerControllerImpl) ReadByUsername(c *gin.Context) {
	customerUsername := c.Param("customer_username")
	if customerUsername != "" {
		c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   web.MessageRes{Message: "Customer username empty!"},
		})
		return
	}

	result, err2 := cci.CustomerUsecase.ReadByUsername(c.Request.Context(), &customerUsername)
	if err2.Error() == "Customer not found!" {
		c.JSON(http.StatusNotFound, web.WebResponse{
			Code:   http.StatusNotFound,
			Status: "NOT FOUND",
			Data:   web.MessageRes{Message: err2.Error()},
		})
		return
	}
	if err2 != nil {
		c.JSON(http.StatusInternalServerError, web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL SERVER ERROR",
			Data:   web.MessageRes{Message: err2.Error()},
		})
		return
	}

	c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   *result,
	})
	return
}

func (cci *CustomerControllerImpl) Update(c *gin.Context) {
	var updateCustomerReq customerReqRes.UpdateCustomerReq

	err := c.ShouldBindJSON(&updateCustomerReq)
	if err.Error() == "Customer not exist!" {
		c.JSON(http.StatusNotFound, web.WebResponse{
			Code:   http.StatusNotFound,
			Status: "NOT FOUND",
			Data:   web.MessageRes{Message: err.Error()},
		})
		return
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   web.MessageRes{Message: err.Error()},
		})
		return
	}

	updateCustomerReq.Username = c.Param("customer_username")
	if updateCustomerReq.Username != "" {
		c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   web.MessageRes{Message: "Customer username empty!"},
		})
		return
	}

	err2 := cci.CustomerUsecase.Update(c.Request.Context(), &updateCustomerReq)
	if err2.Error() == "Customer not found!" {
		c.JSON(http.StatusNotFound, web.WebResponse{
			Code:   http.StatusNotFound,
			Status: "NOT FOUND",
			Data:   web.MessageRes{Message: err2.Error()},
		})
		return
	}
	if err2 != nil {
		c.JSON(http.StatusInternalServerError, web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL SERVER ERROR",
			Data:   web.MessageRes{Message: err2.Error()},
		})
		return
	}

	c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   web.MessageRes{Message: "Success update one customer data!"},
	})
	return
}

func (cci *CustomerControllerImpl) Delete(c *gin.Context) {
	customerUsername := c.Param("customer_username")
	if customerUsername != "" {
		c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   web.MessageRes{Message: "Customer id empty!"},
		})
		return
	}

	err2 := cci.CustomerUsecase.Delete(c.Request.Context(), &customerUsername)
	if err2.Error() == "Customer not found!" {
		c.JSON(http.StatusNotFound, web.WebResponse{
			Code:   http.StatusNotFound,
			Status: "NOT FOUND",
			Data:   web.MessageRes{Message: err2.Error()},
		})
		return
	}

	if err2 != nil {
		c.JSON(http.StatusInternalServerError, web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL SERVER ERROR",
			Data:   web.MessageRes{Message: err2.Error()},
		})
		return
	}

	c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   web.MessageRes{Message: "Success delete customer!"},
	})
	return
}
