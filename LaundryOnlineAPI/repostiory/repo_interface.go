package repostiory

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
