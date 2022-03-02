package usecase

import (
	"LaundryOnlineAPI/model/core"
	"LaundryOnlineAPI/model/web/customerReqRes"
	"LaundryOnlineAPI/repository"
	"context"
	"fmt"
)

type CustomerCRUDUsecaseImpl struct {
	CustomerRepo repository.CustomerRepoInterface
}

func NewCustomerCRUDUsecaseImpl(CustomerRepo repository.CustomerRepoInterface) CustomerCRUDUsecaseInterface {
	return &CustomerCRUDUsecaseImpl{CustomerRepo: CustomerRepo}
}

func (ccui *CustomerCRUDUsecaseImpl) Create(ctx context.Context, req *customerReqRes.CreateCustomerReq) error {
	// Make customer object
	customer := &core.Customer{
		Name:     req.Name,
		Username: req.Username,
		Password: req.Password,
		Address:  req.Address,
	}

	// Call repo to save
	err := ccui.CustomerRepo.Save(ctx, customer)
	if err != nil {
		return err
	}

	return nil
}

func (ccui *CustomerCRUDUsecaseImpl) ReadAll(ctx context.Context) (*[]customerReqRes.ReadCustomerRes, error) {
	// Call repo to get all customer
	customers, err := ccui.CustomerRepo.FindAll(ctx)

	// If error return null object and error
	if err != nil {
		return &[]customerReqRes.ReadCustomerRes{}, err
	}

	// Convert slice customers to slice read customer res
	var responses []customerReqRes.ReadCustomerRes

	for _, customer := range *customers {

		var orders []customerReqRes.CustomerOrder
		for _, order := range customer.Order {
			order := customerReqRes.CustomerOrder{
				Id:         order.ID,
				ServiceId:  order.ServiceID,
				DryWeight:  uint(order.DryWeight),
				TotalPrice: uint(order.TotalPrice),
				Status:     order.Status,
			}
			orders = append(orders, order)
		}

		response := customerReqRes.ReadCustomerRes{
			Name:     customer.Name,
			Username: customer.Username,
			Order:    orders,
			Address:  customer.Address,
		}
		responses = append(responses, response)
	}

	return &responses, nil
}

func (ccui *CustomerCRUDUsecaseImpl) ReadByUsername(ctx context.Context, customerUsername *string) (*customerReqRes.ReadCustomerRes, error) {
	// Call repo to find customer by username
	customer, err := ccui.CustomerRepo.FindByUsername(ctx, customerUsername)
	if err != nil {
		return &customerReqRes.ReadCustomerRes{}, err
	}

	// Convert customer to read customer res
	var orders []customerReqRes.CustomerOrder
	for _, order := range customer.Order {
		order := customerReqRes.CustomerOrder{
			Id:         order.ID,
			ServiceId:  order.ServiceID,
			DryWeight:  uint(order.DryWeight),
			TotalPrice: uint(order.TotalPrice),
			Status:     order.Status,
		}
		orders = append(orders, order)
	}

	response := &customerReqRes.ReadCustomerRes{
		Name:     customer.Name,
		Username: customer.Username,
		Order:    orders,
		Address:  customer.Address,
	}
	return response, nil
}

func (ccui *CustomerCRUDUsecaseImpl) Update(ctx context.Context, req *customerReqRes.UpdateCustomerReq) error {
	// Call repo to check whether customer exist or not
	customer, err := ccui.CustomerRepo.FindByUsername(ctx, &req.Username)
	if err != nil {
		return err
	}

	// Validate whether input is not empty
	if req.Name != "" {
		customer.Name = req.Name
	}
	if req.Username != "" {
		customer.Username = req.Username
		fmt.Println(req.Username)
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

	// Call repo to update customer data
	err2 := ccui.CustomerRepo.Update(ctx, customer)
	if err2 != nil {
		return err2
	}
	return nil
}

func (ccui *CustomerCRUDUsecaseImpl) Delete(ctx context.Context, customerUsername *string) error {
	// Call repo to check whether customer exist or not
	customer, err := ccui.CustomerRepo.FindByUsername(ctx, customerUsername)
	if err != nil {
		return err
	}

	// Call repo to destroy customer data
	err2 := ccui.CustomerRepo.Destroy(ctx, &customer.Username)
	if err2 != nil {
		return err2
	}
	return nil
}
