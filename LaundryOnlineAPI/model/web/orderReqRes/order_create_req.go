package orderReqRes

type CreateOrderReq struct {
	CustomerUsername string `json:"customer_username"`
	ServiceId        uint   `json:"service_id"`
	DryWeight        uint   `json:"dry_weight"`
	TotalPrice       uint32 `json:"total_price"`
	Status           string `json:"status"`
}
