package httphelper

import (
	"encoding/json"
	"net/http"
)

// ErrorResponse todo
func ErrorResponse(w http.ResponseWriter, err error, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write([]byte(`{"error": "` + err.Error() + `"}`))
}

// SuccessResponse todo
func SuccessResponse(w http.ResponseWriter, result any) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(struct {
		Result any `json:"result"`
	}{
		Result: result,
	})
}
