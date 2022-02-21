package orderReqRes

type UpdateOrderReq struct {
	OrderId    uint
	ServiceId  uint   `json:"service_id"`
	DryWeight  uint8  `json:"dry_weight`
	TotalPrice uint32 `json:"total_price"`
	Status     string `json:"status"`
}
