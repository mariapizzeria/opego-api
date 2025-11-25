package order

import (
	"github.com/mariapizzeria/opego-api/pkg/db"
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
