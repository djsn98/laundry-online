package auth

import (
	"LaundryOnlineAPI/model/core"
	"LaundryOnlineAPI/repository"
	"context"
	"errors"
)

func AuthenticateCustomer(ctx context.Context, username string, password string, customerRepo repository.CustomerRepoInterface) (*core.Customer, error) {
	customerInfo, err := customerRepo.FindByUsername(ctx, &username)
	if err != nil {
		return &core.Customer{}, err
	}
	if customerInfo.Name == "" {
		return &core.Customer{}, errors.New("Customer not found!")
	}

	if username == customerInfo.Username {
		if password == customerInfo.Password {
			return customerInfo, nil
		}
		return &core.Customer{}, errors.New("Invalid username or password!")
	}
	return &core.Customer{}, errors.New("Invalid username or password!")
}
