package serviceReqRes

type UpdateServiceReq struct {
	ServiceID  uint
	Name       string `json:"name"`
	PricePerKg uint16 `json:"price_per_kg"`
}
