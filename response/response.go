package response

import (
    "github.com/neotroops/go-realworld/i18n"
    "log"
    "strconv"
)

type Response struct {
    Data    interface{}     `json:"data"`
    Success bool            `json:"success"`
    Errors  []ErrorResponse `json:"errors"`
}

type ErrorResponse struct {
    Code       string `json:"code"`
    Message    string `json:"message"`
    Status     bool   `json:"status"`
    StatusCode int
}

func (e *ErrorResponse) Error() string {
    return e.Message
}

func NewErrorResponse(code string, lang string) ErrorResponse {
    statusCode, err := strconv.Atoi(i18n.ResponseI18nMap[lang][code]["status_code"])
    if err != nil {
        log.Fatal(err)
    }
    return ErrorResponse{
        Message:    i18n.ResponseI18nMap[lang][code]["message"],
        Code:       i18n.ResponseI18nMap[lang][code]["code"],
        StatusCode: statusCode,
        Status:     false,
    }
}
