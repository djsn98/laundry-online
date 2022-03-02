package routes

import (
	"LaundryOnlineAPI/controller"
	"LaundryOnlineAPI/database"
	"LaundryOnlineAPI/repository"
	"LaundryOnlineAPI/usecase"

	"github.com/gin-gonic/gin"
)

// Placeholder to make
var (
	serviceController  controller.ServiceControllerInterface
	customerController controller.CustomerControllerInterface
	orderController    controller.OrderControllerInterface
)

type Route struct {
	URI     string
	Method  string
	Handler func(c *gin.Context)
}

// Merge all route into one route slice
func Load() []Route {
	var routes []Route
	routes = ServiceRoute
	routes = append(routes, OrderRoute...)
	routes = append(routes, CustomerRoute...)

	return routes
}

// Initialized controller for regiter controller method to routes
func init() {
	DB := database.Connection

	//Instance Repo
	serviceRepo := repository.NewServiceRepoImpl(DB)
	customerRepo := repository.NewCustomerRepoImpl(DB)
	orderRepo := repository.NewOrderRepoImpl(DB)

	//Instance Usecase
	serviceUsecase := usecase.NewServiceCRUDUsecaseImpl(serviceRepo)
	customerUsecase := usecase.NewCustomerCRUDUsecaseImpl(customerRepo)
	orderUsecase := usecase.NewOrderCRUDUsecaseImpl(orderRepo, customerRepo)

	//Instance Controller
	serviceController = controller.NewServiceControllerImpl(serviceUsecase)
	customerController = controller.NewCustomerControllerImpl(customerUsecase)
	orderController = controller.NewOrderControllerImpl(orderUsecase)
}
