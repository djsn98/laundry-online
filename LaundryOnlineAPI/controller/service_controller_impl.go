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

	// Get req body and place in variable
	err := c.ShouldBindJSON(&createServiceReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   web.MessageRes{Message: err.Error()},
		})
		return
	}

	// Call service to create Service
	err2 := sci.ServiceUsecase.Create(c.Request.Context(), &createServiceReq)
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
		Data:   web.MessageRes{Message: "Success create new laundry service!"},
	})
	return
}

func (sci *ServiceControllerImpl) ReadAll(c *gin.Context) {
	//  Call service read all service
	result, err := sci.ServiceUsecase.ReadAll(c.Request.Context())

	if err != nil {
		// Error not found handling
		if err.Error() == "Service not found!" || err.Error() == "record not found" {
			c.JSON(http.StatusNotFound, web.WebResponse{
				Code:   http.StatusNotFound,
				Status: "NOT FOUND",
				Data:   web.MessageRes{Message: err.Error()},
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

func (sci *ServiceControllerImpl) ReadById(c *gin.Context) {
	// Get service_id param and convert into integer
	serviceIdInt, err := strconv.Atoi(c.Param("service_id"))

	// If error convert
	if err != nil {
		c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   web.MessageRes{Message: err.Error()},
		})
		return
	}

	// If not input service_id param
	if serviceIdInt == 0 {
		c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   web.MessageRes{Message: "Service id empty!"},
		})
		return
	}

	// Convert Service Id int to uint
	serviceIdUint := uint(serviceIdInt)

	// Call service to read Service by id
	result, err2 := sci.ServiceUsecase.ReadById(c.Request.Context(), &serviceIdUint)

	if err2 != nil {
		// Not found error handling
		if err2.Error() == "Service not found!" || err2.Error() == "record not found" {
			c.JSON(http.StatusNotFound, web.WebResponse{
				Code:   http.StatusNotFound,
				Status: "NOT FOUND",
				Data:   web.MessageRes{Message: "Service not found!"},
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
		Data:   *result,
	})
	return
}

func (sci *ServiceControllerImpl) Update(c *gin.Context) {
	var updateServiceReq serviceReqRes.UpdateServiceReq

	// Get req body and place in variable
	err := c.ShouldBindJSON(&updateServiceReq)

	// If error binding
	if err != nil {
		c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   web.MessageRes{Message: err.Error()},
		})
		return
	}

	// Get service_id param and convert to integer
	serviceIdInt, err := strconv.Atoi(c.Param("service_id"))

	// If error convert
	if err != nil {
		c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   web.MessageRes{Message: err.Error()},
		})
		return
	}

	// If not input service_id param
	if serviceIdInt == 0 {
		c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   web.MessageRes{Message: "Service id empty!"},
		})
		return
	}

	// Convert service_id param into uint
	updateServiceReq.ServiceID = uint(serviceIdInt)

	// Call service update Service
	err2 := sci.ServiceUsecase.Update(c.Request.Context(), &updateServiceReq)

	if err2 != nil {
		// Not found error handling
		if err2.Error() == "Service not found!" || err2.Error() == "record not found" {
			c.JSON(http.StatusNotFound, web.WebResponse{
				Code:   http.StatusNotFound,
				Status: "NOT FOUND",
				Data:   web.MessageRes{Message: "Service not found!"},
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
		Data:   web.MessageRes{Message: "Success update one laundry service!"},
	})
	return
}

func (sci *ServiceControllerImpl) Delete(c *gin.Context) {
	// Get service_id param and convert to integer
	serviceIdInt, err := strconv.Atoi(c.Param("service_id"))

	// If error convert
	if err != nil {
		c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   web.MessageRes{Message: err.Error()},
		})
		return
	}

	// If not input service_id param
	if serviceIdInt == 0 {
		c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   web.MessageRes{Message: "Service id empty!"},
		})
		return
	}

	// Convert service_id param into uint
	serviceIdUint := uint(serviceIdInt)

	// Call service delete Service
	err2 := sci.ServiceUsecase.Delete(c.Request.Context(), &serviceIdUint)

	if err2 != nil {
		// Not found error handling
		if err2.Error() == "Service not found!" || err2.Error() == "record not found" {
			c.JSON(http.StatusNotFound, web.WebResponse{
				Code:   http.StatusNotFound,
				Status: "NOT FOUND",
				Data:   web.MessageRes{Message: "Service not found!"},
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
		Data:   web.MessageRes{Message: "Success delete laundry service!"},
	})
	return
}
