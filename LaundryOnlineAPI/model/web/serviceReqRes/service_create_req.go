package serviceReqRes

type CreateServiceReq struct {
	Name       string `json:"name"`
	PricePerKg uint16 `json:"price_per_kg"`
}
