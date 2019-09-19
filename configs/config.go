package configs

import (
    "github.com/sirupsen/logrus"
    "github.com/spf13/viper"
    "os"
    "sync"
)

type Config struct {
    App App
    Db  Db
}

var (
    config Config
    once   sync.Once
)

func AllConfig() Config {
    viper.SetConfigName("application")
    viper.AddConfigPath("./")
    viper.SetConfigType("yaml")
    if err := viper.ReadInConfig(); err != nil {
        logrus.Panic(err)
        os.Exit(1)
    }

    viper.AutomaticEnv()
    once.Do(func() {
        config = Config{
            App: NewAppConfig(),
            Db:  NewDbConfig(),
        }
    })
    return config
}
