package order

import (
	"time"

	"github.com/mariapizzeria/opego-api/pkg/db"
	"gorm.io/gorm"
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
	res := repo.db.Table("order").Select("order_status").Where("order_id = ? AND canceled_at = NULL", id).Find(&orderStatus)
	if res.Error != nil {
		return "", res.Error
	}
	return orderStatus, nil
}

func (repo *Repository) cancelOrder(id uint) error {
	var now = time.Now()
	res := repo.db.Table("order").Where("order_id = ? AND canceled_at = NULL", id).Update("canceled_at", now)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
