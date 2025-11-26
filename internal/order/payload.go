package order

import (
	"time"

	"github.com/lib/pq"
)

type OrderRequest struct {
	PassengerId      uint           `json:"passenger_id"`
	AddressFrom      string         `json:"address_from"`
	AddressTo        string         `json:"address_to"`
	Tariff           string         `json:"tariff"`
	SelectedServices pq.StringArray `json:"selected_services" gorm:"type:text[]"`
	Comment          string         `json:"comment"`
}

type OrderResponse struct {
	OrderId          uint           `json:"order_id" gorm:"primary_key"`
	UpdatedAt        time.Time      `json:"updated_at"`
	CreatedAt        time.Time      `json:"created_at"`
	CanceledAt       *time.Time     `json:"canceled_at" gorm:"index"`
	PassengerId      uint           `json:"passenger,omitempty" gorm:"foreignKey:PassengerId;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	OrderStatus      string         `json:"order_status"`
	DriverAssigned   *uint          `json:"driver,omitempty" gorm:"default:null"`
	ArrivedCode      string         `json:"arrived_code"`
	AddressFrom      string         `json:"address_from"`
	AddressTo        string         `json:"address_to"`
	Tariff           string         `json:"tariff"`
	SelectedServices pq.StringArray `json:"selected_services" gorm:"type:text[]"`
	Comment          string         `json:"comment"`
	Price            int            `json:"price"`
}

type OrderStatusResponse struct {
	OrderId     uint   `json:"order_id" gorm:"primary_key"`
	OrderStatus string `json:"order_status"`
}
