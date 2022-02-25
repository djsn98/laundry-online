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
	order := &core.Order{
		CustomerUsername: req.CustomerUsername,
		ServiceID:        req.ServiceId,
		DryWeight:        uint8(req.DryWeight),
		TotalPrice:       req.TotalPrice,
		Status:           req.Status,
	}

	err := ocui.OrderRepo.Save(ctx, order)
	if err != nil {
		return err
	}

	return nil
}

func (ocui *OrderCRUDUsecaseImpl) ReadAll(ctx context.Context) (*[]orderReqRes.ReadOrderWithIdRes, error) {
	orders, err := ocui.OrderRepo.FindAll(ctx)
	if err != nil {
		return &[]orderReqRes.ReadOrderWithIdRes{}, err
	}

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
	order, err := ocui.OrderRepo.FindById(ctx, orderId)
	if err != nil {
		return &orderReqRes.ReadOrderRes{}, err
	}

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
	order, err := ocui.OrderRepo.FindById(ctx, &req.OrderId)
	if err != nil {
		return err
	}

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

	err2 := ocui.OrderRepo.Update(ctx, order)
	if err2 != nil {
		return err2
	}
	return nil
}

func (ocui *OrderCRUDUsecaseImpl) Delete(ctx context.Context, orderId *uint) error {
	order, err := ocui.OrderRepo.FindById(ctx, orderId)
	if err != nil {
		return err
	}

	err2 := ocui.OrderRepo.Destroy(ctx, &order.ID)
	if err2 != nil {
		return err2
	}
	return nil
}
