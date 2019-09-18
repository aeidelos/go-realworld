package cmd

import (
    "fmt"
    "github.com/neotroops/go-realworld/constant"
    "github.com/sirupsen/logrus"
    "github.com/spf13/cobra"
    "os"
)

var initCmd = &cobra.Command{
    Use: "gorealworld",
    Short: "Go lang real world implementation",
    Run: func(cmd *cobra.Command, args []string) {
        logrus.Info(fmt.Sprintf(constant.INIT_MESSAGE, 8080))
    },
}

func Exec() {
    if err:= initCmd.Execute(); err != nil {
        logrus.Panic(err)
        os.Exit(1)
    }
}
