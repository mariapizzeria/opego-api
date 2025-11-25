package response

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func JsonEncoder(w http.ResponseWriter, res any, status int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(res)
}

func HandleBody[T any](w http.ResponseWriter, r *http.Request) (*T, error) {
	body, err := decodeBody[T](r.Body)
	if err != nil {
		JsonEncoder(w, err.Error(), http.StatusBadRequest)
		return nil, err
	}
	err = IsValid(body)
	if err != nil {
		JsonEncoder(w, err.Error(), http.StatusBadRequest)
		return nil, err
	}
	return &body, nil

}

func decodeBody[T any](body io.ReadCloser) (T, error) {
	var payload T
	err := json.NewDecoder(body).Decode(&payload)
	if err != nil {
		return payload, err
	}
	return payload, nil
}

func IsValid[T any](body T) error {
	valid := validator.New()
	err := valid.Struct(body)
	return err
}
