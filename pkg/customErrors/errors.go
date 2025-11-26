package customErrors

import "net/http"

func EmptyInput(w http.ResponseWriter) {
	http.Error(w, "Invalid input", http.StatusBadRequest)
}

func ServerError(w http.ResponseWriter) {
	http.Error(w, "Server Error", http.StatusInternalServerError)
}

func CancelOrderError(w http.ResponseWriter) {
	http.Error(w, "Cannot cancel an order. Incorrect status", http.StatusMethodNotAllowed)
}

func OrderNotFoundError(w http.ResponseWriter) {
	http.Error(w, "Order not found or doesn't exist", http.StatusNotFound)
}

func OrderStatusChangedError(w http.ResponseWriter) {
	http.Error(w, "Order status has changed. Order cannot be accepted", http.StatusMethodNotAllowed)
}

func AssignDriverError(w http.ResponseWriter) {
	http.Error(w, "Cannot assign driver", http.StatusMethodNotAllowed)
}
