package usecase

import (
	"LaundryOnlineAPI/auth"
	"LaundryOnlineAPI/config"
	"LaundryOnlineAPI/model/web/customerReqRes"
	"LaundryOnlineAPI/repository"
	"context"
	"errors"
)

type CustomerLoginUsecaseImpl struct {
	CustomerRepo repository.CustomerRepoInterface
}

func NewCustomerLoginUsecaseImpl(CustomerRepo repository.CustomerRepoInterface) CustomerLoginUsecaseInterface {
	return &CustomerLoginUsecaseImpl{CustomerRepo: CustomerRepo}
}

func (clui *CustomerLoginUsecaseImpl) Login(ctx context.Context, req *customerReqRes.LoginCustomerReq) (string, bool, error) {
	var config = config.Config
	customerInfo, err := auth.AuthenticateCustomer(ctx, req.Username, req.Password, clui.CustomerRepo)
	if err != nil {
		return "", false, err
	}

	tokenString, err := auth.GenerateJWT(customerInfo.Username, customerInfo.Password, config.APPLICATION_NAME, config.LOGIN_EXPIRATION_DURATION)
	if err != nil {
		return "", false, err
	}

	if tokenString == "" {
		return "", false, errors.New("Failed generate JWT!")
	}

	return tokenString, true, nil
}
