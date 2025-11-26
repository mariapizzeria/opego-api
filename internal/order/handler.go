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
	statusArrivedCode = "waiting_for_confirmation"
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
	router.HandleFunc("POST /api/order", handler.createOrder())
	router.HandleFunc("POST /api/order/{order_id}/cancel", handler.cancelOrder())
	router.HandleFunc("POST /api/order/{order_id}/accept", handler.acceptOrder())
	router.HandleFunc("POST /api/order/{order_id}/arrived", handler.createArriveCode())
	router.HandleFunc("PUT /api/order/{order_id}/status", handler.updateOrderStatus())
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
		orderId := r.URL.Query().Get("order_id")
		if orderId == "" {
			customErrors.EmptyInput(w)
			return
		}
		// логика отмены заказа
	}
}

func (handler *Handler) acceptOrder() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		orderId := r.URL.Query().Get("order_id")
		if orderId == "" {
			customErrors.EmptyInput(w)
			return
		}
		// логика принятия заказа исполнителем
	}
}

func (handler *Handler) createArriveCode() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		orderIdStr := r.URL.Query().Get("order_id")
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
			OrderStatus: statusArrivedCode,
			ArrivedCode: code,
		}
		response.JsonEncoder(w, result, 200)
	}
}

func (handler *Handler) updateOrderStatus() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		orderId := r.URL.Query().Get("order_id")
		if orderId == "" {
			customErrors.EmptyInput(w)
			return
		}
		// Обновление статуса заказа
	}
}
