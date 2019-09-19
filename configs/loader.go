package configs

import (
    "fmt"
    "github.com/sirupsen/logrus"
    "github.com/spf13/viper"
    "os"
)

func GetStringConfiguration(key string) string {
    checkKeyAvailable(key)
    return viper.GetString(key)
}

func GetIntConfiguration(key string) int {
    checkKeyAvailable(key)
    return viper.GetInt(key)
}

func checkKeyAvailable(key string) {
    if !viper.IsSet(key) {
        logrus.Panic(fmt.Sprintf("keys %s is not set", key))
        os.Exit(1)
    }
}
