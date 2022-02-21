package customerReqRes

import "LaundryOnlineAPI/model/core"

type ReadCustomerRes struct {
	Name     string        `json:"name"`
	Username string        `json:"username"`
	Order    *[]core.Order `json:"order"`
	Address  string        `json:"address"`
}
