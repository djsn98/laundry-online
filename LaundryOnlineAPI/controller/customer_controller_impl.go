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

	// Get req body and place into variable
	err := c.ShouldBindJSON(&createCustomerReq)

	// If error send BAD REQUEST response and error mn
	if err != nil {
		c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   web.MessageRes{Message: err.Error()},
		})
		return
	}

	// Call service to create Customer
	err2 := cci.CustomerUsecase.Create(c.Request.Context(), &createCustomerReq)

	// If error send INTERNAL SERVER ERROR response and error message
	if err2 != nil {
		c.JSON(http.StatusInternalServerError, web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL SERVER ERROR",
			Data:   web.MessageRes{Message: err2.Error()},
		})
		return
	}

	// If success send OK response
	c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   web.MessageRes{Message: "Success add new customer!"},
	})
	return
}

func (cci *CustomerControllerImpl) ReadAll(c *gin.Context) {
	// Call service read all customer
	result, err := cci.CustomerUsecase.ReadAll(c.Request.Context())

	// If error
	if err != nil {
		// Error not found handling
		if err.Error() == "Customer not found!" || err.Error() == "record not found" {
			c.JSON(http.StatusNotFound, web.WebResponse{
				Code:   http.StatusNotFound,
				Status: "NOT FOUND",
				Data:   web.MessageRes{Message: "Customer not found!"},
			})
			return
		}

		// Generic error handling
		c.JSON(http.StatusInternalServerError, web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL SERVER ERROR",
			Data:   web.MessageRes{Message: err.Error()},
		})
		return
	}

	// If success send OK response
	c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   *result,
	})
	return
}

func (cci *CustomerControllerImpl) ReadByUsername(c *gin.Context) {
	// get customer_username param
	customerUsername := c.Param("customer_username")

	// If not input customer_username param
	if customerUsername == "" {
		c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   web.MessageRes{Message: "Customer username empty!"},
		})
		return
	}

	// Call service read customer by username
	result, err2 := cci.CustomerUsecase.ReadByUsername(c.Request.Context(), &customerUsername)
	if err2 != nil {
		// Error not foun handling
		if err2.Error() == "Customer not found!" || err2.Error() == "record not found" {
			c.JSON(http.StatusNotFound, web.WebResponse{
				Code:   http.StatusNotFound,
				Status: "NOT FOUND",
				Data:   web.MessageRes{Message: "Customer not found!"},
			})
			return
		}
		// Generic error handling
		c.JSON(http.StatusInternalServerError, web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL SERVER ERROR",
			Data:   web.MessageRes{Message: err2.Error()},
		})
		return
	}

	// If success send OK response
	c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   *result,
	})
	return
}

func (cci *CustomerControllerImpl) Update(c *gin.Context) {
	var updateCustomerReq customerReqRes.UpdateCustomerReq

	// Get req body and place into variable
	err := c.ShouldBindJSON(&updateCustomerReq)
	// If error send BAD REQUEST response
	if err != nil {
		c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   web.MessageRes{Message: err.Error()},
		})
		return
	}

	// Get customer_username param
	updateCustomerReq.Username = c.Param("customer_username")

	// If not input customer_username param
	if updateCustomerReq.Username == "" {
		c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   web.MessageRes{Message: "Customer username empty!"},
		})
		return
	}

	// Call service update customer
	err2 := cci.CustomerUsecase.Update(c.Request.Context(), &updateCustomerReq)
	if err2 != nil {
		// Not found error handling
		if err2.Error() == "Customer not found!" || err2.Error() == "record not found" {
			c.JSON(http.StatusNotFound, web.WebResponse{
				Code:   http.StatusNotFound,
				Status: "NOT FOUND",
				Data:   web.MessageRes{Message: "Customer not found!"},
			})
			return
		}

		// Generic error handling
		c.JSON(http.StatusInternalServerError, web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL SERVER ERROR",
			Data:   web.MessageRes{Message: err2.Error()},
		})
		return
	}

	// If success send OK response
	c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   web.MessageRes{Message: "Success update one customer data!"},
	})
	return
}

func (cci *CustomerControllerImpl) Delete(c *gin.Context) {
	// Get customer_username param
	customerUsername := c.Param("customer_username")

	// If not input customer_username param
	if customerUsername == "" {
		c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   web.MessageRes{Message: "Customer id empty!"},
		})
		return
	}

	// Call service delete customer
	err2 := cci.CustomerUsecase.Delete(c.Request.Context(), &customerUsername)
	if err2 != nil {
		// Not found error handling
		if err2.Error() == "Customer not found!" || err2.Error() == "record not found" {
			c.JSON(http.StatusNotFound, web.WebResponse{
				Code:   http.StatusNotFound,
				Status: "NOT FOUND",
				Data:   web.MessageRes{Message: "Customer not found!"},
			})
			return
		}

		// Generic error handling
		c.JSON(http.StatusInternalServerError, web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL SERVER ERROR",
			Data:   web.MessageRes{Message: err2.Error()},
		})
		return
	}

	// If success send OK response
	c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   web.MessageRes{Message: "Success delete customer!"},
	})
	return
}
