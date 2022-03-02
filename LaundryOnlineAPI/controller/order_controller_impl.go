package controller

import (
	"LaundryOnlineAPI/model/web"
	"LaundryOnlineAPI/model/web/orderReqRes"
	"LaundryOnlineAPI/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrderControllerImpl struct {
	OrderUsecase usecase.OrderCRUDUsecaseInterface
}

func NewOrderControllerImpl(OrderUsecase usecase.OrderCRUDUsecaseInterface) OrderControllerInterface {
	return &OrderControllerImpl{OrderUsecase: OrderUsecase}
}

func (oci *OrderControllerImpl) Create(c *gin.Context) {
	var createOrderReq orderReqRes.CreateOrderReq

	// Get req body and place into variable
	err := c.ShouldBindJSON(&createOrderReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   web.MessageRes{Message: err.Error()},
		})
		return
	}

	err2 := oci.OrderUsecase.Create(c.Request.Context(), &createOrderReq)
	if err2 != nil {
		c.JSON(http.StatusInternalServerError, web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL SERVER ERROR",
			Data:   web.MessageRes{Message: err2.Error()},
		})
		return
	}

	// If success
	c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   web.MessageRes{Message: "Success create new order!"},
	})
	return
}

func (oci *OrderControllerImpl) ReadAll(c *gin.Context) {
	// Call service read all order
	result, err := oci.OrderUsecase.ReadAll(c.Request.Context())
	if err != nil {
		// Not found error handling
		if err.Error() == "Order not found!" || err.Error() == "record not found" {
			c.JSON(http.StatusNotFound, web.WebResponse{
				Code:   http.StatusNotFound,
				Status: "NOT FOUND",
				Data:   web.MessageRes{Message: "Order not found!"},
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

	// If success
	c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   *result,
	})
	return
}

func (oci *OrderControllerImpl) ReadById(c *gin.Context) {
	// Get order_id param and convert to integer
	orderIdInt, err := strconv.Atoi(c.Param("order_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   web.MessageRes{Message: err.Error()},
		})
		return
	}

	// If not input order_id
	if orderIdInt == 0 {
		c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   web.MessageRes{Message: "Order id empty!"},
		})
		return
	}

	// Convert order_id from int to uint
	orderIdUint := uint(orderIdInt)

	// Call service read order by id
	result, err2 := oci.OrderUsecase.ReadById(c.Request.Context(), &orderIdUint)
	if err2 != nil {
		// Not found error handling
		if err2.Error() == "Order not found!" || err2.Error() == "record not found" {
			c.JSON(http.StatusNotFound, web.WebResponse{
				Code:   http.StatusNotFound,
				Status: "NOT FOUND",
				Data:   web.MessageRes{Message: "Order not found!"},
			})
			return
		}
		// Generic error handling service
		c.JSON(http.StatusInternalServerError, web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL SERVER ERROR",
			Data:   web.MessageRes{Message: err2.Error()},
		})
		return
	}

	// If success
	c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   *result,
	})
	return
}

func (oci *OrderControllerImpl) Update(c *gin.Context) {
	var updateOrderReq orderReqRes.UpdateOrderReq

	// Get req body and place into variable
	err := c.ShouldBindJSON(&updateOrderReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   web.MessageRes{Message: err.Error()},
		})
		return
	}

	// Get order_id param
	orderIdInt, err := strconv.Atoi(c.Param("order_id"))

	// If error convert
	if err != nil {
		c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   web.MessageRes{Message: err.Error()},
		})
		return
	}

	// If not input order_id
	if orderIdInt == 0 {
		c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   web.MessageRes{Message: "Order id empty!"},
		})
		return
	}

	// Convert order_id and plaace in updateOrderReq atribute
	updateOrderReq.OrderId = uint(orderIdInt)
	err2 := oci.OrderUsecase.Update(c.Request.Context(), &updateOrderReq)
	if err2 != nil {
		// Not found error handling
		if err2.Error() == "Order not found!" || err2.Error() == "record not found" {
			c.JSON(http.StatusNotFound, web.WebResponse{
				Code:   http.StatusNotFound,
				Status: "NOT FOUND",
				Data:   web.MessageRes{Message: "Order not found!"},
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

	// If success
	c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   web.MessageRes{Message: "Success update one order!"},
	})
	return
}

func (oci *OrderControllerImpl) Delete(c *gin.Context) {
	// Get order_id param and convert to int
	orderIdInt, err := strconv.Atoi(c.Param("order_id"))

	// If error convert
	if err != nil {
		c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   web.MessageRes{Message: err.Error()},
		})
		return
	}

	// If not input order_id param
	if orderIdInt == 0 {
		c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   web.MessageRes{Message: "Order id empty!"},
		})
		return
	}

	// Convert order_id from int to uint
	orderIdUint := uint(orderIdInt)

	// Call service delete order
	err2 := oci.OrderUsecase.Delete(c.Request.Context(), &orderIdUint)
	if err2 != nil {
		// Not found error handling
		if err2.Error() == "Order not found!" || err2.Error() == "record not found" {
			c.JSON(http.StatusNotFound, web.WebResponse{
				Code:   http.StatusNotFound,
				Status: "NOT FOUND",
				Data:   web.MessageRes{Message: "Order not found!"},
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

	// If success
	c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   web.MessageRes{Message: "Success delete order!"},
	})
	return
}
