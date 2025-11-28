package customErrors

import (
	"net/http"

	"github.com/mariapizzeria/opego-api/pkg/response"
)

type CustomErrors struct {
	Error string `json:"error"`
}

func EmptyInput(w http.ResponseWriter) {
	response.JsonEncoder(w, "Input cannot be empty", 400)
}

func ParseDataError(w http.ResponseWriter, err error) {
	response.JsonEncoder(w, CustomErrors{
		Error: err.Error(),
	}, 422)
}

func ServerError(w http.ResponseWriter, err error) {
	response.JsonEncoder(w, CustomErrors{
		Error: err.Error(),
	}, 500)
}

func CancelOrderError(w http.ResponseWriter) {
	response.JsonEncoder(w, "Cannot cancel an order. Incorrect status", 422)
}

func OrderStatusChangedError(w http.ResponseWriter) {
	response.JsonEncoder(w, "Order status has changed", 422)
}

func AssignDriverError(w http.ResponseWriter, err error) {
	response.JsonEncoder(w, CustomErrors{
		Error: err.Error(),
	}, 422)
}

func DriverIsNotAvailable(w http.ResponseWriter, err error) {
	response.JsonEncoder(w, CustomErrors{
		Error: err.Error(),
	}, 422)
}

func CalculationError(w http.ResponseWriter, err error) {
	response.JsonEncoder(w, CustomErrors{
		Error: err.Error(),
	}, 422)
}

func CreateRecordError(w http.ResponseWriter, err error) {
	response.JsonEncoder(w, CustomErrors{
		Error: err.Error(),
	}, 422)
}

func UpdateRecordError(w http.ResponseWriter, err error) {
	response.JsonEncoder(w, CustomErrors{
		Error: err.Error(),
	}, 422)
}

func GetRecordError(w http.ResponseWriter, err error) {
	response.JsonEncoder(w, CustomErrors{
		Error: err.Error(),
	}, 422)
}
