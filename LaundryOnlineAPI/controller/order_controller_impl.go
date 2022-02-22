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

func (oci *OrderControllerImpl) Create(c *gin.Context) {
	var createOrderReq orderReqRes.CreateOrderReq

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

	c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   web.MessageRes{Message: "Success create new order!"},
	})
	return
}

func (oci *OrderControllerImpl) ReadAll(c *gin.Context) {
	result, err := oci.OrderUsecase.ReadAll(c.Request.Context())
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

func (oci *OrderControllerImpl) ReadById(c *gin.Context) {
	orderIdInt, err := strconv.Atoi(c.Param("order_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   web.MessageRes{Message: err.Error()},
		})
		return
	}
	if orderIdInt == 0 {
		c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   web.MessageRes{Message: "Order id empty!"},
		})
		return
	}

	orderIdUint := uint(orderIdInt)
	result, err2 := oci.OrderUsecase.ReadById(c.Request.Context(), &orderIdUint)
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

func (oci *OrderControllerImpl) Update(c *gin.Context) {
	var updateOrderReq orderReqRes.UpdateOrderReq

	err := c.ShouldBindJSON(&updateOrderReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   web.MessageRes{Message: err.Error()},
		})
		return
	}

	orderIdInt, err := strconv.Atoi(c.Param("order_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   web.MessageRes{Message: err.Error()},
		})
		return
	}
	if orderIdInt == 0 {
		c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   web.MessageRes{Message: "Order id empty!"},
		})
		return
	}

	updateOrderReq.OrderId = uint(orderIdInt)
	err2 := oci.OrderUsecase.Update(c.Request.Context(), &updateOrderReq)
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
		Data:   web.MessageRes{Message: "Success update one order!"},
	})
	return
}

func (oci *OrderControllerImpl) Delete(c *gin.Context) {
	orderIdInt, err := strconv.Atoi(c.Param("order_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   web.MessageRes{Message: err.Error()},
		})
		return
	}
	if orderIdInt == 0 {
		c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   web.MessageRes{Message: "Order id empty!"},
		})
		return
	}

	orderIdUint := uint(orderIdInt)
	err2 := oci.OrderUsecase.Delete(c.Request.Context(), &orderIdUint)
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
		Data:   web.MessageRes{Message: "Success delete order!"},
	})
	return
}
