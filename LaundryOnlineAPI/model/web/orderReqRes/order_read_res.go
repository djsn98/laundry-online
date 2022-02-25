package orderReqRes

type ReadOrderRes struct {
	CustomerUsername string `json:"customer_username"`
	ServiceId        uint   `json:"service_id"`
	DryWeight        uint8  `json:"dry_weight"`
	TotalPrice       uint32 `json:"total_price"`
	Status           string `json:"status"`
}

type ReadOrderWithIdRes struct {
	Id               uint   `json:"id"`
	CustomerUsername string `json:"customer_username"`
	ServiceId        uint   `json:"service_id"`
	DryWeight        uint8  `json:"dry_weight"`
	TotalPrice       uint32 `json:"total_price"`
	Status           string `json:"status"`
}
