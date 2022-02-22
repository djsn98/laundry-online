package usecase

import (
	"LaundryOnlineAPI/model/web/orderReqRes"
	"context"
)

type OrderCRUDUsecaseInterface interface {
	Create(ctx context.Context, req *orderReqRes.CreateOrderReq) error
	ReadAll(ctx context.Context) (*[]orderReqRes.ReadOrderRes, error)
	ReadById(ctx context.Context, orderId *uint) (*orderReqRes.ReadOrderRes, error)
	Update(ctx context.Context, req *orderReqRes.UpdateOrderReq) error
	Delete(ctx context.Context, orderId *uint) error
}
