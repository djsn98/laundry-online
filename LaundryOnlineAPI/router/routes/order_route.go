package routes

import "net/http"

var OrderRoute = []Route{
	{
		URI:     "/order",
		Method:  http.MethodPost,
		Handler: orderController.Create,
	},
	{
		URI:     "/orders",
		Method:  http.MethodGet,
		Handler: orderController.ReadAll,
	},
	{
		URI:     "/order/:order_id",
		Method:  http.MethodGet,
		Handler: orderController.ReadById,
	},
	{
		URI:     "/order/:order_id",
		Method:  http.MethodPut,
		Handler: orderController.Update,
	},
	{
		URI:     "/order/:order_id",
		Method:  http.MethodDelete,
		Handler: orderController.Delete,
	},
}
