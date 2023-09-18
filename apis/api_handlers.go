package apis

import (
    "net/http"
    "encoding/json"
)

func MyAPIHandler(w http.ResponseWriter, r *http.Request) {
    data := map[string]string{"message": "This is an API response"}
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(data)
}
