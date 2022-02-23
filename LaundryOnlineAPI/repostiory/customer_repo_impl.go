package repostiory

import (
	"LaundryOnlineAPI/model/core"
	"context"
	"errors"

	"gorm.io/gorm"
)

type CustomerRepoImpl struct {
	DB *gorm.DB
}

func NewCustomerRepoImpl(DB *gorm.DB) CustomerRepoInterface {
	return &CustomerRepoImpl{DB: DB}
}

func (cri *CustomerRepoImpl) Save(ctx context.Context, customer *core.Customer) error {
	err := cri.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(customer).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (cri *CustomerRepoImpl) FindAll(ctx context.Context) (*[]core.Customer, error) {
	var customers []core.Customer

	err := cri.DB.Find(&customers).Error
	if err != nil {
		return &customers, err
	}

	if len(customers) == 0 {
		return &customers, errors.New("Customer not found!")
	}

	return &customers, nil
}

func (cri *CustomerRepoImpl) FindByUsername(ctx context.Context, customerUsername *string) (*core.Customer, error) {
	var customer core.Customer

	err := cri.DB.First(&customer, customerUsername).Error
	if err != nil {
		return &customer, err
	}

	if customer.Name == "" {
		return &customer, errors.New("Customer not found!")
	}

	return &customer, nil
}

func (cri *CustomerRepoImpl) Update(ctx context.Context, customer *core.Customer) error {
	err := cri.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(customer).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (cri *CustomerRepoImpl) Destroy(ctx context.Context, customerUsername *string) error {
	err := cri.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&core.Customer{}, customerUsername).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
