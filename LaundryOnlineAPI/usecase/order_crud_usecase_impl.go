package usecase

import (
	"LaundryOnlineAPI/model/core"
	"LaundryOnlineAPI/model/web/orderReqRes"
	"LaundryOnlineAPI/repostiory"
	"context"
)

type OrderCRUDUsecaseImpl struct {
	OrderRepo repostiory.OrderRepoInterface
}

func (ocui *OrderCRUDUsecaseImpl) Create(ctx context.Context, req *orderReqRes.CreateOrderReq) error {
	order := &core.Order{
		CustomerID: req.CustomerId,
		ServiceID:  req.ServiceId,
		DryWeight:  uint8(req.DryWeight),
		TotalPrice: req.TotalPrice,
		Status:     req.Status,
	}

	err := ocui.OrderRepo.Save(ctx, order)
	if err != nil {
		return err
	}

	return nil
}

func (ocui *OrderCRUDUsecaseImpl) ReadAll(ctx context.Context) (*[]orderReqRes.ReadOrderRes, error) {
	orders, err := ocui.OrderRepo.FindAll(ctx)
	if err != nil {
		return &[]orderReqRes.ReadOrderRes{}, err
	}

	var responses []orderReqRes.ReadOrderRes
	for _, order := range *orders {
		response := orderReqRes.ReadOrderRes{
			CustomerId: order.CustomerID,
			ServiceId:  order.ServiceID,
			DryWeight:  order.DryWeight,
			TotalPrice: order.TotalPrice,
			Status:     order.Status,
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
		CustomerId: order.CustomerID,
		ServiceId:  order.ServiceID,
		DryWeight:  order.DryWeight,
		TotalPrice: order.TotalPrice,
		Status:     order.Status,
	}
	return response, nil
}

func (ocui *OrderCRUDUsecaseImpl) Update(ctx context.Context, req *orderReqRes.UpdateOrderReq) error {
	order := &core.Order{
		ServiceID:  req.ServiceId,
		DryWeight:  req.DryWeight,
		TotalPrice: req.TotalPrice,
		Status:     req.Status,
	}
	order.ID = req.OrderId

	err := ocui.OrderRepo.Update(ctx, order)
	if err != nil {
		return err
	}
	return nil
}

func (ocui *OrderCRUDUsecaseImpl) Delete(ctx context.Context, orderId *uint) error {
	err := ocui.OrderRepo.Destroy(ctx, orderId)
	if err != nil {
		return err
	}
	return nil
}
