package customerReqRes

type UpdateCustomerReq struct {
	Name     string `json:"name"`
	Username string
	Password string `json:"password"`
	Balance  uint32 `json:"balance"`
	Address  string `json:"address"`
}
