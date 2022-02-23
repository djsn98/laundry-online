package usecase

import (
	"LaundryOnlineAPI/model/core"
	"LaundryOnlineAPI/model/web/customerReqRes"
	"LaundryOnlineAPI/repository"
	"context"
)

type CustomerCRUDUsecaseImpl struct {
	CustomerRepo repository.CustomerRepoInterface
}

func NewCustomerCRUDUsecaseImpl(CustomerRepo repository.CustomerRepoInterface) CustomerCRUDUsecaseInterface {
	return &CustomerCRUDUsecaseImpl{CustomerRepo: CustomerRepo}
}

func (ccui *CustomerCRUDUsecaseImpl) Create(ctx context.Context, req *customerReqRes.CreateCustomerReq) error {
	customer := &core.Customer{
		Name:     req.Name,
		Username: req.Username,
		Password: req.Password,
		Address:  req.Address,
	}

	err := ccui.CustomerRepo.Save(ctx, customer)
	if err != nil {
		return err
	}

	return nil
}

func (ccui *CustomerCRUDUsecaseImpl) ReadAll(ctx context.Context) (*[]customerReqRes.ReadCustomerRes, error) {
	customers, err := ccui.CustomerRepo.FindAll(ctx)
	if err != nil {
		return &[]customerReqRes.ReadCustomerRes{}, err
	}

	var responses []customerReqRes.ReadCustomerRes
	for _, customer := range *customers {
		response := customerReqRes.ReadCustomerRes{
			Name:     customer.Name,
			Username: customer.Username,
			Order:    &customer.Order,
			Address:  customer.Address,
		}
		responses = append(responses, response)
	}

	return &responses, nil
}

func (ccui *CustomerCRUDUsecaseImpl) ReadByUsername(ctx context.Context, customerUsername *string) (*customerReqRes.ReadCustomerRes, error) {
	customer, err := ccui.CustomerRepo.FindByUsername(ctx, customerUsername)
	if err != nil {
		return &customerReqRes.ReadCustomerRes{}, err
	}

	response := &customerReqRes.ReadCustomerRes{
		Name:     customer.Name,
		Username: customer.Username,
		Order:    &customer.Order,
		Address:  customer.Address,
	}
	return response, nil
}

func (ccui *CustomerCRUDUsecaseImpl) Update(ctx context.Context, req *customerReqRes.UpdateCustomerReq) error {
	customer, err := ccui.CustomerRepo.FindByUsername(ctx, &req.Username)
	if err != nil {
		return err
	}

	if req.Name != "" {
		customer.Name = req.Name
	}
	if req.Username != "" {
		customer.Username = req.Username
	}
	if req.Password != "" {
		customer.Password = req.Password
	}
	if req.Balance != 0 {
		customer.Balance = req.Balance
	}
	if req.Address != "" {
		customer.Address = req.Address
	}

	err2 := ccui.CustomerRepo.Update(ctx, customer)
	if err2 != nil {
		return err2
	}
	return nil
}

func (ccui *CustomerCRUDUsecaseImpl) Delete(ctx context.Context, customerUsername *string) error {
	customer, err := ccui.CustomerRepo.FindByUsername(ctx, customerUsername)
	if err != nil {
		return err
	}

	err2 := ccui.CustomerRepo.Destroy(ctx, &customer.Username)
	if err2 != nil {
		return err2
	}
	return nil
}
