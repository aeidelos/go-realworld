package cmd

import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    "github.com/golang-migrate/migrate"
    "github.com/golang-migrate/migrate/database/mysql"
    _ "github.com/golang-migrate/migrate/source/file"
    "github.com/neotroops/go-realworld/configs"
    "github.com/neotroops/go-realworld/constants"
    "github.com/neotroops/go-realworld/i18n"
    "github.com/neotroops/go-realworld/pkg"
    "github.com/sirupsen/logrus"
    "github.com/spf13/cobra"
    "os"
    "os/signal"
    "syscall"
)

var initCmd = &cobra.Command{
    Use:   "gorealworld",
    Short: "Go lang real world implementation",
    Run: func(cmd *cobra.Command, args []string) {
        go serverShutdown()
        config := configs.AllConfig().App
        i18n.Init()
        if config.AppEnv == "dev" {
            logrus.SetReportCaller(true)
            logrus.SetLevel(logrus.DebugLevel)
        }
        logrus.Info(fmt.Sprintf(constants.InitMessage, config.AppName, config.AppPort))
        pkg.StartAPIServer(config)
    },
}

func serverShutdown() {
    s := make(chan os.Signal)
    signal.Notify(s, os.Interrupt, syscall.SIGTERM)
    go func() {
        <- s
        defer os.Exit(0)
        logrus.Info("shutdown server")
    }()
}

var migrateCmd = &cobra.Command{
    Use:   "migrate",
    Short: "Go lang real world implementation",
    Run: func(cmd *cobra.Command, args []string) {
        config := configs.AllConfig().Db
        dbUser := config.DbUser
        dbPassword := config.DbPassword
        dbHost := config.DbHost
        dbPort := config.DbPort
        dbName := config.DbName
        db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?multiStatements=true", dbUser, dbPassword, dbHost, dbPort, dbName))
        if err != nil {
            logrus.Panic(err)
            os.Exit(1)
        }
        driver, _ := mysql.WithInstance(db, &mysql.Config{})
        m, _ := migrate.NewWithDatabaseInstance(
            "file://db/migrations",
            "mysql",
            driver,
        )
        if err := m.Steps(1); err != nil {
            logrus.Panic(err)
            os.Exit(1)
        }

    },
}

func Exec() {
    initCmd.AddCommand(migrateCmd)
    if err := initCmd.Execute(); err != nil {
        logrus.Panic(err)
        os.Exit(1)
    }
}
