package response

import (
    "github.com/neotroops/go-realworld/configs"
    "github.com/neotroops/go-realworld/i18n"
    "log"
    "strconv"
)


type ErrorResponse struct {
    Code       int    `json:"code"`
    Message    string `json:"message"`
    Status     bool   `json:"status"`
    StatusCode int    `json:"status_code"`
}

func (e *ErrorResponse) Error() string {
    return e.Message
}

func NewErrorResponse(key string, status bool, config configs.App) ErrorResponse {
    lang := config.AppLang
    statusCode, err := strconv.Atoi(i18n.ResponseI18nMap[lang][key]["status_code"])
    if err != nil {
        log.Fatal(err)
    }
    code, err := strconv.Atoi(i18n.ResponseI18nMap[lang][key]["code"])
    if err != nil {
        log.Fatal(err)
    }
    return ErrorResponse{
        Message:    i18n.ResponseI18nMap[lang][key]["message"],
        Code:       code,
        StatusCode: statusCode,
        Status:     status,
    }
}
