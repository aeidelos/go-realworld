package configs

type App struct {
    AppName string
    AppPort string
    AppLang string
    AppEnv  string
}

func NewAppConfig() App {
    return App{
        AppName: GetStringConfiguration("APP_NAME"),
        AppPort: GetStringConfiguration("APP_PORT"),
        AppLang: GetStringConfiguration("APP_LANG"),
        AppEnv:  GetStringConfiguration("APP_ENV"),
    }
}
