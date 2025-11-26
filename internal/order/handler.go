package order

import (
	"net/http"
	"strconv"

	"github.com/mariapizzeria/opego-api/pkg/customErrors"
	"github.com/mariapizzeria/opego-api/pkg/response"
	"github.com/mariapizzeria/opego-api/services/notifications"
	"github.com/mariapizzeria/opego-api/services/priceCalculator"
)

const (
	orderStatusWaitingForConfirmation = "waiting_for_confirmation"
	orderStatusInProgress             = "in_progress"
	orderStatusCompleted              = "completed"
	orderStatusSearching              = "searching"
	orderStatusDriverAssigned         = "driver_assigned"
	orderStatusCanceled               = "canceled"
)

type Handler struct {
	Repository Repository
}

type HandlerDeps struct {
	Repository Repository
}

func NewHandler(router *http.ServeMux, deps HandlerDeps) {
	handler := Handler{
		deps.Repository,
	}
	router.HandleFunc("GET /api/order/{order_id}", handler.getOrderStatus())
	router.HandleFunc("POST /api/order", handler.createOrder())                   // сделано
	router.HandleFunc("POST /api/order/{order_id}/cancel", handler.cancelOrder()) // сделано
	router.HandleFunc("POST /api/order/{order_id}/accept", handler.acceptOrder()) // частично сделано. Добавить проверку кук пользователя
	router.HandleFunc("POST /api/driver/status", handler.acceptDriverStatus())    // сделано
	router.HandleFunc("POST /api/order/{order_id}/arrived", handler.createArriveCode())
	router.HandleFunc("PUT /api/order/{order_id}/status", handler.updateOrderStatus())
	router.HandleFunc("PUT /api/order/{order_id}/status/search", handler.updateOrderStatusSearch()) // сделано
}

func (handler *Handler) getOrderStatus() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		orderId := r.URL.Query().Get("order_id")
		if orderId == "" {
			customErrors.EmptyInput(w)
			return
		}
		// Далее логика получение статуса в реальном времени
	}
}

func (handler *Handler) createOrder() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := response.HandleBody[OrderRequest](w, r)
		if err != nil {
			customErrors.ServerError(w)
			return
		}
		price, err := priceCalculator.PriceCalculation(
			body.Tariff,
			body.AddressFrom,
			body.AddressTo,
			body.SelectedServices,
		)
		if err != nil {
			customErrors.ServerError(w)
			return
		}
		result, err := handler.Repository.createOrder(&OrderResponse{
			PassengerId:      body.PassengerId,
			OrderStatus:      "pending",
			AddressFrom:      body.AddressFrom,
			AddressTo:        body.AddressTo,
			Tariff:           body.Tariff,
			SelectedServices: body.SelectedServices,
			Comment:          body.Comment,
			Price:            price,
		})
		if err != nil {
			customErrors.ServerError(w)
			return
		}
		response.JsonEncoder(w, result, 201)

	}
}

func (handler *Handler) cancelOrder() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		orderIdStr := r.PathValue("order_id")
		if orderIdStr == "" {
			customErrors.EmptyInput(w)
			return
		}
		idUint, err := strconv.ParseUint(orderIdStr, 10, 64)
		if err != nil {
			customErrors.ServerError(w)
			return
		}
		orderStatus, err := handler.Repository.getOrderStatus(uint(idUint))
		if err != nil {
			customErrors.ServerError(w)
			return
		}
		if orderStatus == orderStatusCompleted || orderStatus == orderStatusInProgress {
			customErrors.CancelOrderError(w)
			return
		}
		err = handler.Repository.cancelOrder(uint(idUint))
		if err != nil {
			customErrors.OrderNotFoundError(w)
			return
		}
		_, err = handler.Repository.updateOrderStatus(&OrderStatusResponse{
			OrderId:     uint(idUint),
			OrderStatus: orderStatusCanceled,
		})
		response.JsonEncoder(w, nil, 204)
	}
}

func (handler *Handler) acceptDriverStatus() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := response.HandleBody[Driver](w, r)
		if err != nil {
			customErrors.ServerError(w)
			return
		}
		res, err := handler.Repository.updateDriverStatus(&DriverStatus{
			DriverId:        body.DriverId,
			Available:       body.Available,
			CurrentLocation: body.CurrentLocation,
		})
		if err != nil {
			customErrors.DriverIsNotAvailable(w)
			return
		}
		response.JsonEncoder(w, res, 200)
	}
}

func (handler *Handler) acceptOrder() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		orderIdStr := r.PathValue("order_id")
		if orderIdStr == "" {
			customErrors.EmptyInput(w)
			return
		}
		orderId, err := strconv.Atoi(orderIdStr)
		if err != nil {
			customErrors.ServerError(w)
			return
		}
		orderStatus, err := handler.Repository.getOrderStatus(uint(orderId))
		if err != nil {
			customErrors.OrderNotFoundError(w)
			return
		}
		if orderStatus != orderStatusSearching {
			customErrors.OrderStatusChangedError(w)
			return
		}

		//здесь должна быть проверка на session_id
		// если id сесси прошло проверку то изменяем статус заказа передавая id

		var session_id uint
		session_id = 2

		driver, err := handler.Repository.assignDriver(uint(orderId), session_id)
		if err != nil {
			customErrors.AssignDriverError(w)
			return
		}
		_, err = handler.Repository.updateOrderStatus(&OrderStatusResponse{
			OrderId:     uint(orderId),
			OrderStatus: orderStatusDriverAssigned,
		})
		if err != nil {
			customErrors.OrderNotFoundError(w)
			return
		}
		response.JsonEncoder(w, driver, 200)
	}
}

func (handler *Handler) createArriveCode() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		orderIdStr := r.PathValue("order_id")
		if orderIdStr == "" {
			customErrors.EmptyInput(w)
			return
		}
		orderId, err := strconv.Atoi(orderIdStr)
		if err != nil {
			customErrors.ServerError(w)
			return
		}
		code := notifications.GenerateArrivedCode()
		result := OrderResponse{
			OrderId:     uint(orderId),
			OrderStatus: orderStatusWaitingForConfirmation,
			ArrivedCode: code,
		}
		response.JsonEncoder(w, result, 200)
	}
}

func (handler *Handler) updateOrderStatus() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		orderId := r.PathValue("order_id")
		if orderId == "" {
			customErrors.EmptyInput(w)
			return
		}
		// получение заказа, если статус pending -> search
		// если водитель подтвердил заказ, то статус заказа меняется с  search а driver_assigned + id водителя
		// Водитель подъехал - ждет код. Статус меняется на waiting_for_confirmation
		// Водитель подтвердил начало поездки - статус in_progress
		// Водитель подтвердил конец поездки - статус completed
		// Обновление статуса заказа
	}
}

// Обновление пользователем статуса заказа при нажатии на кнопку "Заказать"
func (handler *Handler) updateOrderStatusSearch() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		orderIdStr := r.PathValue("order_id")
		if orderIdStr == "" {
			customErrors.EmptyInput(w)
			return
		}
		orderId, err := strconv.Atoi(orderIdStr)
		if err != nil {
			customErrors.ServerError(w)
			return
		}
		OrderStatus, err := handler.Repository.updateOrderStatus(&OrderStatusResponse{
			OrderId:     uint(orderId),
			OrderStatus: orderStatusSearching,
		})
		if err != nil {
			customErrors.OrderNotFoundError(w)
			return
		}
		response.JsonEncoder(w, OrderStatus, 201)
	}
}
