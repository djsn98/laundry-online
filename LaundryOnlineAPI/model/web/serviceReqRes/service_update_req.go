package serviceReqRes

type UpdateServiceReq struct {
	ServiceID  uint   `json:"service_id"`
	Name       string `json:"name"`
	PricePerKg uint16 `json:"price_per_kg"`
}
