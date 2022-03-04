package auth

import (
	"LaundryOnlineAPI/repository"
	"context"
	"errors"
)

func AuthenticateCustomer(ctx context.Context, username string, password string, customerRepo repository.CustomerRepoInterface) (bool, error) {
	customerInfo, err := customerRepo.FindByUsername(ctx, &username)
	if err != nil {
		return false, err
	}
	if customerInfo.Name == "" {
		return false, errors.New("Customer not found!")
	}

	if username == customerInfo.Username {
		if password == customerInfo.Password {
			return true, nil
		}
		return false, errors.New("Invalid username or password!")
	}
	return false, errors.New("Invalid username or password!")
}
