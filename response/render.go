package response

import (
    "encoding/json"
    "github.com/sirupsen/logrus"
    "net/http"
)

func RenderErrorResponse(w http.ResponseWriter, response ErrorResponse) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(response.StatusCode)
    err := json.NewEncoder(w).Encode(response)
    if err != nil {
        logrus.Fatal(err)
    }
}
