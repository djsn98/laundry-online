package usecase

import (
	"LaundryOnlineAPI/model/core"
	"LaundryOnlineAPI/model/web/orderReqRes"
	"LaundryOnlineAPI/repository"
	"context"
)

type OrderCRUDUsecaseImpl struct {
	OrderRepo    repository.OrderRepoInterface
	CustomerRepo repository.CustomerRepoInterface
}

func NewOrderCRUDUsecaseImpl(OrderRepo repository.OrderRepoInterface, CustomerRepo repository.CustomerRepoInterface) OrderCRUDUsecaseInterface {
	return &OrderCRUDUsecaseImpl{OrderRepo: OrderRepo, CustomerRepo: CustomerRepo}
}

func (ocui *OrderCRUDUsecaseImpl) Create(ctx context.Context, req *orderReqRes.CreateOrderReq) error {
	// Create new object order and input req data
	order := &core.Order{
		CustomerUsername: req.CustomerUsername,
		ServiceID:        req.ServiceId,
		DryWeight:        uint8(req.DryWeight),
		TotalPrice:       req.TotalPrice,
		Status:           req.Status,
	}

	// Calll repo to save data
	err := ocui.OrderRepo.Save(ctx, order)
	if err != nil {
		return err
	}

	return nil
}

func (ocui *OrderCRUDUsecaseImpl) ReadAll(ctx context.Context) (*[]orderReqRes.ReadOrderWithIdRes, error) {
	// Call repo to get all ordee
	orders, err := ocui.OrderRepo.FindAll(ctx)

	// If error return null object and error
	if err != nil {
		return &[]orderReqRes.ReadOrderWithIdRes{}, err
	}

	// Convert slice order into slice read order with id res
	var responses []orderReqRes.ReadOrderWithIdRes
	for _, order := range *orders {
		response := orderReqRes.ReadOrderWithIdRes{
			Id:               order.ID,
			CustomerUsername: order.CustomerUsername,
			ServiceId:        order.ServiceID,
			DryWeight:        order.DryWeight,
			TotalPrice:       order.TotalPrice,
			Status:           order.Status,
		}
		responses = append(responses, response)
	}

	return &responses, nil
}

func (ocui *OrderCRUDUsecaseImpl) ReadById(ctx context.Context, orderId *uint) (*orderReqRes.ReadOrderRes, error) {
	// call repo to find order by id
	order, err := ocui.OrderRepo.FindById(ctx, orderId)
	if err != nil {
		return &orderReqRes.ReadOrderRes{}, err
	}

	// Convert order to read order res
	response := &orderReqRes.ReadOrderRes{
		CustomerUsername: order.CustomerUsername,
		ServiceId:        order.ServiceID,
		DryWeight:        order.DryWeight,
		TotalPrice:       order.TotalPrice,
		Status:           order.Status,
	}
	return response, nil
}

func (ocui *OrderCRUDUsecaseImpl) Update(ctx context.Context, req *orderReqRes.UpdateOrderReq) error {
	// Call repo to check whether order exist or not
	order, err := ocui.OrderRepo.FindById(ctx, &req.OrderId)
	if err != nil {
		return err
	}

	// Check whether req empty or not
	if req.ServiceId != 0 {
		order.ServiceID = req.ServiceId
	}
	if req.DryWeight != 0 {
		order.DryWeight = uint8(req.DryWeight)
	}
	if req.TotalPrice != 0 {
		order.TotalPrice = uint32(req.TotalPrice)
	}
	if req.Status != "" {
		order.Status = req.Status
	}

	// Call repo to update order
	err2 := ocui.OrderRepo.Update(ctx, order)
	if err2 != nil {
		return err2
	}
	return nil
}

func (ocui *OrderCRUDUsecaseImpl) Delete(ctx context.Context, orderId *uint) error {
	// Call repo to check whether order exist or not
	order, err := ocui.OrderRepo.FindById(ctx, orderId)
	if err != nil {
		return err
	}

	// Call repo to destroy order data
	err2 := ocui.OrderRepo.Destroy(ctx, &order.ID)
	if err2 != nil {
		return err2
	}
	return nil
}
