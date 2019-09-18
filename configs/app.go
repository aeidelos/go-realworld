package configs

import (
    "github.com/sirupsen/logrus"
    "github.com/spf13/viper"
    "os"
    "sync"
)

type AppConfig struct {
    AppName  string
    AppPort  string
    AppLang  string
    DbConfig DbConfig
}

var (
    appConfig AppConfig
    once      sync.Once
)

func Config() AppConfig {
    viper.SetConfigName("application")
    viper.AddConfigPath("./")
    viper.SetConfigType("yaml")
    if err := viper.ReadInConfig(); err != nil {
        logrus.Panic(err)
        os.Exit(1)
    }

    viper.AutomaticEnv()
    once.Do(func() {
        appConfig = AppConfig{
            AppName: GetStringConfiguration("APP_NAME"),
            AppPort: GetStringConfiguration("APP_PORT"),
            AppLang: GetStringConfiguration("APP_LANG"),
            DbConfig: NewDbConfig(),
        }
    })
    return appConfig
}
