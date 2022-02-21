package customerReqRes

type CreateCustomerReq struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Address  string `json:"address"`
}
