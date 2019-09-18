package cmd

import (
    "fmt"
    "github.com/neotroops/go-realworld/configs"
    "github.com/neotroops/go-realworld/constant"
    "github.com/sirupsen/logrus"
    "github.com/spf13/cobra"
    "os"
)

var initCmd = &cobra.Command{
    Use: "gorealworld",
    Short: "Go lang real world implementation",
    Run: func(cmd *cobra.Command, args []string) {
        appConfig := configs.Config()
        logrus.Info(fmt.Sprintf(constant.INIT_MESSAGE, appConfig.AppName, appConfig.AppPort))
    },
}

func Exec() {
    if err:= initCmd.Execute(); err != nil {
        logrus.Panic(err)
        os.Exit(1)
    }
}
