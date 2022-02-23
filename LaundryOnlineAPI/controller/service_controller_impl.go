package controller

import (
	"LaundryOnlineAPI/model/web"
	"LaundryOnlineAPI/model/web/serviceReqRes"
	"LaundryOnlineAPI/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ServiceControllerImpl struct {
	ServiceUsecase usecase.ServiceCRUDUsecaseInterface
}

func NewServiceControllerImpl(ServiceUsecase usecase.ServiceCRUDUsecaseInterface) ServiceControllerInterface {
	return &ServiceControllerImpl{ServiceUsecase: ServiceUsecase}
}

func (sci *ServiceControllerImpl) Create(c *gin.Context) {
	var createServiceReq serviceReqRes.CreateServiceReq

	err := c.ShouldBindJSON(&createServiceReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   web.MessageRes{Message: err.Error()},
		})
		return
	}

	err2 := sci.ServiceUsecase.Create(c.Request.Context(), &createServiceReq)
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
		Data:   web.MessageRes{Message: "Success create new laundry service!"},
	})
	return
}

func (sci *ServiceControllerImpl) ReadAll(c *gin.Context) {
	result, err := sci.ServiceUsecase.ReadAll(c.Request.Context())
	if err != nil {
		if err.Error() == "Service not found!" {
			c.JSON(http.StatusNotFound, web.WebResponse{
				Code:   http.StatusNotFound,
				Status: "NOT FOUND",
				Data:   web.MessageRes{Message: err.Error()},
			})
			return
		}

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

func (sci *ServiceControllerImpl) ReadById(c *gin.Context) {
	serviceIdInt, err := strconv.Atoi(c.Param("service_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   web.MessageRes{Message: err.Error()},
		})
		return
	}
	if serviceIdInt == 0 {
		c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   web.MessageRes{Message: "Service id empty!"},
		})
		return
	}

	serviceIdUint := uint(serviceIdInt)
	result, err2 := sci.ServiceUsecase.ReadById(c.Request.Context(), &serviceIdUint)
	if err2 != nil {
		if err2.Error() == "Service not found!" {
			c.JSON(http.StatusNotFound, web.WebResponse{
				Code:   http.StatusNotFound,
				Status: "NOT FOUND",
				Data:   web.MessageRes{Message: err2.Error()},
			})
			return
		}
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

func (sci *ServiceControllerImpl) Update(c *gin.Context) {
	var updateServiceReq serviceReqRes.UpdateServiceReq

	err := c.ShouldBindJSON(&updateServiceReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   web.MessageRes{Message: err.Error()},
		})
		return
	}

	serviceIdInt, err := strconv.Atoi(c.Param("service_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   web.MessageRes{Message: err.Error()},
		})
		return
	}
	if serviceIdInt == 0 {
		c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   web.MessageRes{Message: "Service id empty!"},
		})
		return
	}

	updateServiceReq.ServiceID = uint(serviceIdInt)
	err2 := sci.ServiceUsecase.Update(c.Request.Context(), &updateServiceReq)
	if err2 != nil {
		if err2.Error() == "Service not found!" {
			c.JSON(http.StatusNotFound, web.WebResponse{
				Code:   http.StatusNotFound,
				Status: "NOT FOUND",
				Data:   web.MessageRes{Message: err2.Error()},
			})
			return
		}
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
		Data:   web.MessageRes{Message: "Success update one laundry service!"},
	})
	return
}

func (sci *ServiceControllerImpl) Delete(c *gin.Context) {
	serviceIdInt, err := strconv.Atoi(c.Param("service_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   web.MessageRes{Message: err.Error()},
		})
		return
	}
	if serviceIdInt == 0 {
		c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   web.MessageRes{Message: "Service id empty!"},
		})
		return
	}

	serviceIdUint := uint(serviceIdInt)
	err2 := sci.ServiceUsecase.Delete(c.Request.Context(), &serviceIdUint)
	if err2 != nil {
		if err2.Error() == "Service not found!" {
			c.JSON(http.StatusNotFound, web.WebResponse{
				Code:   http.StatusNotFound,
				Status: "NOT FOUND",
				Data:   web.MessageRes{Message: err2.Error()},
			})
			return
		}
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
		Data:   web.MessageRes{Message: "Success delete laundry service!"},
	})
	return
}
