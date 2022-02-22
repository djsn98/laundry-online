package usecase

import (
	"LaundryOnlineAPI/model/web/customerReqRes"
	"LaundryOnlineAPI/model/web/orderReqRes"
	"LaundryOnlineAPI/model/web/serviceReqRes"
	"context"
)

type OrderCRUDUsecaseInterface interface {
	Create(ctx context.Context, req *orderReqRes.CreateOrderReq) error
	ReadAll(ctx context.Context) (*[]orderReqRes.ReadOrderRes, error)
	ReadById(ctx context.Context, orderId *uint) (*orderReqRes.ReadOrderRes, error)
	Update(ctx context.Context, req *orderReqRes.UpdateOrderReq) error
	Delete(ctx context.Context, orderId *uint) error
}

type ServiceCRUDUsecaseInterface interface {
	Create(ctx context.Context, req *serviceReqRes.CreateServiceReq) error
	ReadAll(ctx context.Context) (*[]serviceReqRes.ReadServiceRes, error)
	ReadById(ctx context.Context, serviceId *uint) (*serviceReqRes.ReadServiceRes, error)
	Update(ctx context.Context, req *serviceReqRes.UpdateServiceReq) error
	Delete(ctx context.Context, serviceId *uint) error
}

type CustomerCRUDUsecaseInterface interface {
	Create(ctx context.Context, req *customerReqRes.CreateCustomerReq) error
	ReadAll(ctx context.Context) (*[]customerReqRes.ReadCustomerRes, error)
	ReadByUsername(ctx context.Context, customerUsername *string) (*customerReqRes.ReadCustomerRes, error)
	Update(ctx context.Context, req *customerReqRes.UpdateCustomerReq) error
	Delete(ctx context.Context, customerUsername *string) error
}
