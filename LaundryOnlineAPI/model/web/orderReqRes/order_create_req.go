package orderReqRes

type CreateOrderReq struct {
	CustomerId uint   `json:"customer_id"`
	ServiceId  uint   `json:"service_id"`
	DryWeight  uint   `json:"dry_weight"`
	TotalPrice uint32 `json:"total_price"`
	Status     string `json:"status"`
}
