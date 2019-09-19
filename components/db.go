package components

import (
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    "github.com/jmoiron/sqlx"
    "github.com/neotroops/go-realworld/configs"
    "github.com/sirupsen/logrus"
    "sync"
    "time"
)

var once sync.Once

var connection *sqlx.DB

func GetConnection(config configs.Db) *sqlx.DB {
    once.Do(func() {
        conn, err := sqlx.Connect(config.DbDriver,
            fmt.Sprintf("%s:%s@tcp(%s:%v)/%s",
                config.DbUser,
                config.DbPassword,
                config.DbHost,
                config.DbPort,
                config.DbName))
        if err != nil {
            logrus.Panic(err)
            panic("db connection can't be made")
        }
        if conn.Ping() != nil {
            panic("ping failed to start")
        }
        conn.SetConnMaxLifetime(time.Duration(config.DbMaxLifeTimeConnection))
        conn.SetMaxIdleConns(config.DbMaxIdleConnection)
        conn.SetMaxOpenConns(config.DbMaxConnection)
        connection = conn
    })
    return connection
}
