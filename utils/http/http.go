package http_utils

import (
	"encoding/json"
	"net/http"

	errors_utils "github.com/bRRRITSCOLD/bookstore_utils-go/errors"
)

func RespondJson(w http.ResponseWriter, statusCode int, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(body)
}

func RespondError(w http.ResponseWriter, err errors_utils.APIError) {
	RespondJson(w, err.Status(), err)
}
