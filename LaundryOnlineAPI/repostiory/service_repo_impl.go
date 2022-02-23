package repostiory

import (
	"LaundryOnlineAPI/model/core"
	"context"
	"errors"

	"gorm.io/gorm"
)

type ServiceRepoImpl struct {
	DB *gorm.DB
}

func NewServiceRepoImpl(DB *gorm.DB) ServiceRepoInterface {
	return &ServiceRepoImpl{DB: DB}
}

func (sri *ServiceRepoImpl) Save(ctx context.Context, service *core.Service) error {
	err := sri.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(service).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (sri *ServiceRepoImpl) FindAll(ctx context.Context) (*[]core.Service, error) {
	var services []core.Service

	err := sri.DB.Find(&services).Error
	if err != nil {
		return &services, err
	}

	if len(services) == 0 {
		return &services, errors.New("Service not found!")
	}

	return &services, nil
}

func (sri *ServiceRepoImpl) FindById(ctx context.Context, serviceId *uint) (*core.Service, error) {
	var service core.Service

	err := sri.DB.First(&service, serviceId).Error
	if err != nil {
		return &service, err
	}

	if service == (core.Service{}) {
		return &service, errors.New("Service not found!")
	}

	return &service, nil
}

func (sri *ServiceRepoImpl) Update(ctx context.Context, service *core.Service) error {
	err := sri.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(service).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (sri *ServiceRepoImpl) Destroy(ctx context.Context, serviceId *uint) error {
	err := sri.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&core.Service{}, serviceId).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
