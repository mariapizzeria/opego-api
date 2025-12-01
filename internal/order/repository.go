package order

import (
	"time"

	"github.com/mariapizzeria/opego-api/pkg/db"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository struct {
	db *db.Db
}

func NewRepository(database *db.Db) Repository {
	return Repository{database}
}

func (repo *Repository) createOrder(order *OrderResponse) (*OrderResponse, error) {
	res := repo.db.Table("order").Create(order)
	if res.Error != nil {
		return nil, res.Error
	}
	return order, nil
}

func (repo *Repository) getOrderStatus(id uint) (string, error) {
	var orderStatus string
	res := repo.db.Table("order").Select("order_status").Where("order_id = ? AND canceled_at IS NULL", id).Find(&orderStatus)
	if res.Error != nil {
		return "", res.Error
	}
	return orderStatus, nil
}

func (repo *Repository) cancelOrder(id uint) error {
	var now = time.Now()
	res := repo.db.Table("order").Where("order_id = ? AND canceled_at IS NULL", id).Update("canceled_at", now)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	upt := repo.db.Table("order").Where("order_id = ?", id).Update("updated_at", now)
	if upt.Error != nil {
		return upt.Error
	}
	return nil
}

func (repo *Repository) updateOrderStatus(order *OrderStatusResponse) (*OrderStatusResponse, error) {
	now := time.Now()
	res := repo.db.Table("order").Where("order_id =?", order.OrderId).Update("order_status", order.OrderStatus)
	if res.Error != nil {
		return nil, res.Error
	}
	if order.OrderStatus == orderStatusCompleted {
		r := repo.db.Table("order").Where("order_id =?", order.OrderId).Update("completed_at", now)
		if r.Error != nil {
			return nil, r.Error
		}
	}
	upt := repo.db.Table("order").Where("order_id = ?", order.OrderId).Update("updated_at", now)
	if upt.Error != nil {
		return nil, upt.Error
	}
	return order, nil
}

func (repo *Repository) getDriverById(id uint) (*Driver, error) {
	var driver *Driver
	err := repo.db.Table("driver").Where("driver_id = ? AND available", id).First(&driver).Error
	if err != nil {
		return nil, err
	}
	return driver, nil
}

func (repo *Repository) assignDriver(orderId uint, driver *Driver) (*Driver, error) {
	res := repo.db.Table("order").Where("order_id = ? AND driver_assigned IS NULL AND canceled_at IS NULL", orderId).Update("driver_assigned", driver.DriverId)
	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	now := time.Now()
	upt := repo.db.Table("order").Where("order_id = ?", orderId).Update("updated_at", now)
	if upt.Error != nil {
		return nil, upt.Error
	}
	return driver, nil
}

func (repo *Repository) updateDriverStatus(driver *DriverStatus) (*DriverStatus, error) {
	res := repo.db.Table("driver").Where("driver_id =?", driver.DriverId).Select("available", "current_location").Updates(&driver)
	if res.Error != nil {
		return nil, res.Error
	}
	return driver, nil
}

func (repo *Repository) createConfirmationCode(order *ConfirmationCode) (*ConfirmationCode, error) {
	res := repo.db.Table("order").Clauses(clause.Returning{}).Updates(&order)
	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	now := time.Now()
	upt := repo.db.Table("order").Where("order_id = ?", order.OrderId).Update("updated_at", now)
	if upt.Error != nil {
		return nil, upt.Error
	}
	return order, nil
}
