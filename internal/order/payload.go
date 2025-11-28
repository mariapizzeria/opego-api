package order

import (
	"time"

	"github.com/lib/pq"
)

type OrderRequest struct {
	PassengerId      uint           `json:"passenger_id" validate:"required"`
	AddressFrom      string         `json:"address_from" validate:"required"`
	AddressTo        string         `json:"address_to" validate:"required"`
	Tariff           string         `json:"tariff" validate:"required,oneof=basic comfort comfort_plus business"`
	SelectedServices pq.StringArray `json:"selected_services" gorm:"type:text[]" validate:"required"`
	Comment          string         `json:"comment"`
}

type OrderResponse struct {
	OrderId          uint           `json:"order_id" gorm:"primary_key"`
	UpdatedAt        time.Time      `json:"updated_at"`
	CreatedAt        time.Time      `json:"created_at"`
	CanceledAt       *time.Time     `json:"canceled_at" gorm:"index"`
	CompletedAt      *time.Time     `json:"completed_at" gorm:"index"`
	PassengerId      uint           `json:"passenger,omitempty" gorm:"foreignKey:PassengerId;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT" validate:"required"`
	OrderStatus      string         `json:"order_status"`
	DriverAssigned   *uint          `json:"driver,omitempty" gorm:"default:null"`
	ArrivedCode      string         `json:"arrived_code"`
	AddressFrom      string         `json:"address_from" validate:"required"`
	AddressTo        string         `json:"address_to" validate:"required"`
	Tariff           string         `json:"tariff" validate:"required,oneof=basic comfort comfort_plus business"`
	SelectedServices pq.StringArray `json:"selected_services" gorm:"type:text[]" validate:"required"`
	Comment          string         `json:"comment"`
	Price            int            `json:"price" validate:"required"`
}

type OrderStatusResponse struct {
	OrderId     uint   `json:"order_id" gorm:"primary_key"`
	OrderStatus string `json:"order_status" validate:"required"`
}

type Driver struct {
	DriverId        uint           `json:"driver_id" gorm:"primary_key"`
	Name            string         `json:"name"`
	CarType         string         `json:"car_type"`
	CarNumber       string         `json:"car_number"`
	Score           string         `json:"score"`
	Available       bool           `json:"available"`
	CurrentLocation DriverLocation `json:"current_location" gorm:"type:jsonb" validate:"required"`
}

type DriverStatus struct {
	DriverId        uint           `json:"driver_id" gorm:"primary_key"`
	Available       bool           `json:"available" validate:"required"`
	CurrentLocation DriverLocation `json:"current_location" gorm:"type:jsonb" validate:"required"`
}

type DriverLocation struct {
	Lat float64 `json:"lat"`
	Ing float64 `json:"ing"`
}

type ConfirmationCode struct {
	OrderId     uint   `json:"order_id" gorm:"primary_key"`
	OrderStatus string `json:"order_status" validate:"required"`
	ArrivedCode string `json:"arrived_code"`
}
