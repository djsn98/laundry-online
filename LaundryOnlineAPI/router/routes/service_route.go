package routes

import "net/http"

var ServiceRoute = []Route{
	{
		URI:     "/service",
		Method:  http.MethodPost,
		Handler: serviceController.Create,
	},
	{
		URI:     "/services",
		Method:  http.MethodGet,
		Handler: serviceController.ReadAll,
	},
	{
		URI:     "/service/:service_id",
		Method:  http.MethodGet,
		Handler: serviceController.ReadById,
	},
	{
		URI:     "/service/:service_id",
		Method:  http.MethodPut,
		Handler: serviceController.Update,
	},
	{
		URI:     "/service/:service_id",
		Method:  http.MethodDelete,
		Handler: serviceController.Delete,
	},
}
