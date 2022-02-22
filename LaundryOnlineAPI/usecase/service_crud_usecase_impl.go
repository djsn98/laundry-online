package usecase

import (
	"LaundryOnlineAPI/model/core"
	"LaundryOnlineAPI/model/web/serviceReqRes"
	"LaundryOnlineAPI/repostiory"
	"context"
	"errors"
)

type ServiceCRUDUsecaseImpl struct {
	ServiceRepo repostiory.ServiceRepoInterface
}

func (scui *ServiceCRUDUsecaseImpl) Create(ctx context.Context, req *serviceReqRes.CreateServiceReq) error {
	service := &core.Service{
		Name:       req.Name,
		PricePerKg: req.PricePerKg,
	}

	err := scui.ServiceRepo.Save(ctx, service)
	if err != nil {
		return err
	}

	return nil
}

func (scui *ServiceCRUDUsecaseImpl) ReadAll(ctx context.Context) (*[]serviceReqRes.ReadServiceRes, error) {
	services, err := scui.ServiceRepo.FindAll(ctx)
	if err != nil {
		return &[]serviceReqRes.ReadServiceRes{}, err
	}

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
	service, err := scui.ServiceRepo.FindById(ctx, serviceId)
	if err != nil {
		return &serviceReqRes.ReadServiceRes{}, err
	}

	response := &serviceReqRes.ReadServiceRes{
		ID:         service.ID,
		Name:       service.Name,
		PricePerKg: service.PricePerKg,
	}
	return response, nil
}

func (scui *ServiceCRUDUsecaseImpl) Update(ctx context.Context, req *serviceReqRes.UpdateServiceReq) error {
	service, err := scui.ServiceRepo.FindById(ctx, &req.ServiceID)
	if err != nil {
		return err
	}

	if service.ID == 0 {
		return errors.New("Order not exist!")
	}

	if req.Name != "" {
		service.Name = req.Name
	}
	if req.PricePerKg != 0 {
		service.PricePerKg = req.PricePerKg
	}

	err2 := scui.ServiceRepo.Update(ctx, service)
	if err2 != nil {
		return err2
	}
	return nil
}

func (scui *ServiceCRUDUsecaseImpl) Delete(ctx context.Context, serviceId *uint) error {
	service, err := scui.ServiceRepo.FindById(ctx, serviceId)
	if err != nil {
		return err
	}

	if service.ID == 0 {
		return errors.New("Order not exist!")
	}

	err2 := scui.ServiceRepo.Destroy(ctx, &service.ID)
	if err2 != nil {
		return err2
	}
	return nil
}
