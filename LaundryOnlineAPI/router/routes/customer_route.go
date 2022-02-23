package routes

import "net/http"

var CustomerRoute = []Route{
	{
		URI:     "/customer",
		Method:  http.MethodPost,
		Handler: customerController.Create,
	},
	{
		URI:     "/customers",
		Method:  http.MethodGet,
		Handler: customerController.ReadAll,
	},
	{
		URI:     "/customer/:customer_username",
		Method:  http.MethodGet,
		Handler: customerController.ReadByUsername,
	},
	{
		URI:     "/customer/:customer_username",
		Method:  http.MethodPut,
		Handler: customerController.Update,
	},
	{
		URI:     "/customer/:customer_username",
		Method:  http.MethodDelete,
		Handler: customerController.Delete,
	},
}
