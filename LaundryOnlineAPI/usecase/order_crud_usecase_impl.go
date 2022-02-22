package usecase

import (
	"LaundryOnlineAPI/model/core"
	"LaundryOnlineAPI/model/web/orderReqRes"
	"LaundryOnlineAPI/repostiory"
	"context"
	"errors"
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
	order, err := ocui.OrderRepo.FindById(ctx, &req.OrderId)
	if err != nil {
		return err
	}

	if order.ID == 0 {
		return errors.New("Order not exist!")
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

	if order.ID == 0 {
		return errors.New("Order not exist!")
	}

	err2 := ocui.OrderRepo.Destroy(ctx, &order.ID)
	if err2 != nil {
		return err2
	}
	return nil
}
