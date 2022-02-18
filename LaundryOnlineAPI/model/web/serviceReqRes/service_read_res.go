package serviceReqRes

type ReadServiceRes struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	PricePerKg uint16 `json:"price_per_kg"`
}
