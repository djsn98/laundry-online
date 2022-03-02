package usecase

import (
	"LaundryOnlineAPI/model/core"
	"LaundryOnlineAPI/model/web/serviceReqRes"
	"LaundryOnlineAPI/repository"
	"context"
)

type ServiceCRUDUsecaseImpl struct {
	ServiceRepo repository.ServiceRepoInterface
}

func NewServiceCRUDUsecaseImpl(ServiceRepo repository.ServiceRepoInterface) ServiceCRUDUsecaseInterface {
	return &ServiceCRUDUsecaseImpl{ServiceRepo: ServiceRepo}
}

func (scui *ServiceCRUDUsecaseImpl) Create(ctx context.Context, req *serviceReqRes.CreateServiceReq) error {
	// Create new object service and input req data
	service := &core.Service{
		Name:       req.Name,
		PricePerKg: req.PricePerKg,
	}

	// Call repo to save new service
	err := scui.ServiceRepo.Save(ctx, service)
	if err != nil {
		return err
	}

	return nil
}

func (scui *ServiceCRUDUsecaseImpl) ReadAll(ctx context.Context) (*[]serviceReqRes.ReadServiceRes, error) {
	// Call repo to get all service
	services, err := scui.ServiceRepo.FindAll(ctx)
	if err != nil {
		return &[]serviceReqRes.ReadServiceRes{}, err
	}

	// Convert slice services to slice read service res
	var responses []serviceReqRes.ReadServiceRes
	for _, service := range *services {
		response := serviceReqRes.ReadServiceRes{
			ID:         service.ID,
			Name:       service.Name,
			PricePerKg: service.PricePerKg,
		}
		responses = append(responses, response)
	}

	return &responses, nil
}

func (scui *ServiceCRUDUsecaseImpl) ReadById(ctx context.Context, serviceId *uint) (*serviceReqRes.ReadServiceRes, error) {
	// Call repo to get service by id
	service, err := scui.ServiceRepo.FindById(ctx, serviceId)
	if err != nil {
		return &serviceReqRes.ReadServiceRes{}, err
	}

	// Convert service to read service response
	response := &serviceReqRes.ReadServiceRes{
		ID:         service.ID,
		Name:       service.Name,
		PricePerKg: service.PricePerKg,
	}
	return response, nil
}

func (scui *ServiceCRUDUsecaseImpl) Update(ctx context.Context, req *serviceReqRes.UpdateServiceReq) error {
	// Call repo to check whether service exist or not
	service, err := scui.ServiceRepo.FindById(ctx, &req.ServiceID)
	if err != nil {
		return err
	}
	// Check whether req empty or not
	if req.Name != "" {
		service.Name = req.Name
	}
	if req.PricePerKg != 0 {
		service.PricePerKg = req.PricePerKg
	}

	// Call repo to update service data
	err2 := scui.ServiceRepo.Update(ctx, service)
	if err2 != nil {
		return err2
	}
	return nil
}

func (scui *ServiceCRUDUsecaseImpl) Delete(ctx context.Context, serviceId *uint) error {
	// Call repo to check whether service exist or not
	service, err := scui.ServiceRepo.FindById(ctx, serviceId)
	if err != nil {
		return err
	}

	// Call repo to destroy service data
	err2 := scui.ServiceRepo.Destroy(ctx, &service.ID)
	if err2 != nil {
		return err2
	}
	return nil
}
