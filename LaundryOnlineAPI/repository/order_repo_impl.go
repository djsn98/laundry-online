package repository

import (
	"LaundryOnlineAPI/model/core"
	"context"
	"errors"

	"gorm.io/gorm"
)

type OrderRepoImpl struct {
	DB *gorm.DB
}

func NewOrderRepoImpl(DB *gorm.DB) OrderRepoInterface {
	return &OrderRepoImpl{DB: DB}
}

func (ori *OrderRepoImpl) Save(ctx context.Context, order *core.Order) error {
	err := ori.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(order).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (ori *OrderRepoImpl) FindAll(ctx context.Context) (*[]core.Order, error) {
	var orders []core.Order

	err := ori.DB.Find(&orders).Error
	if err != nil {
		return &orders, err
	}

	if len(orders) == 0 {
		return &orders, errors.New("Order not found!")
	}

	return &orders, nil
}

func (ori *OrderRepoImpl) FindById(ctx context.Context, orderId *uint) (*core.Order, error) {
	var order core.Order

	err := ori.DB.First(&order, orderId).Error
	if err != nil {
		return &order, err
	}

	if order.Status == "" {
		return &order, errors.New("Order not found!")
	}

	return &order, nil
}

func (ori *OrderRepoImpl) Update(ctx context.Context, order *core.Order) error {
	err := ori.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(order).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (ori *OrderRepoImpl) Destroy(ctx context.Context, orderId *uint) error {
	err := ori.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&core.Order{}, orderId).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
