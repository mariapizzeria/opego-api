package customErrors

import "net/http"

func EmptyInput(w http.ResponseWriter) {
	http.Error(w, "Invalid input", http.StatusBadRequest)
}

func ServerError(w http.ResponseWriter) {
	http.Error(w, "Server Error", http.StatusInternalServerError)
}
