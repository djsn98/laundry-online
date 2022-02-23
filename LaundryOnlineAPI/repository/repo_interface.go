package repository

import (
	"LaundryOnlineAPI/model/core"
	"context"
)

type ServiceRepoInterface interface {
	Save(ctx context.Context, service *core.Service) error
	FindAll(ctx context.Context) (*[]core.Service, error)
	FindById(ctx context.Context, serviceId *uint) (*core.Service, error)
	Update(ctx context.Context, service *core.Service) error
	Destroy(ctx context.Context, serviceId *uint) error
}

type CustomerRepoInterface interface {
	Save(ctx context.Context, customer *core.Customer) error
	FindAll(ctx context.Context) (*[]core.Customer, error)
	FindByUsername(ctx context.Context, customerUsername *string) (*core.Customer, error)
	Update(ctx context.Context, customer *core.Customer) error
	Destroy(ctx context.Context, customerUsername *string) error
}

type OrderRepoInterface interface {
	Save(ctx context.Context, order *core.Order) error
	FindAll(ctx context.Context) (*[]core.Order, error)
	FindById(ctx context.Context, orderId *uint) (*core.Order, error)
	Update(ctx context.Context, order *core.Order) error
	Destroy(ctx context.Context, orderId *uint) error
}
