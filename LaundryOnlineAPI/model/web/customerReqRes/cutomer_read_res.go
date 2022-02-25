package customerReqRes

type ReadCustomerRes struct {
	Name     string          `json:"name"`
	Username string          `json:"username"`
	Order    []CustomerOrder `json:"order"`
	Address  string          `json:"address"`
}

type CustomerOrder struct {
	Id         uint   `json:"id"`
	ServiceId  uint   `json:"service_id"`
	DryWeight  uint   `json:"dry_weight"`
	TotalPrice uint   `json:"total_price"`
	Status     string `json:"status"`
}
