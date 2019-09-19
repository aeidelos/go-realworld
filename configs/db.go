package configs

type Db struct {
    DbUser                  string
    DbPassword              string
    DbHost                  string
    DbPort                  string
    DbName                  string
    DbDriver                string
    DbMaxConnection         int
    DbMaxIdleConnection     int
    DbMaxLifeTimeConnection int
}

func NewDbConfig() Db {
    return Db{
        DbUser:                  GetStringConfiguration("DB_USER"),
        DbPassword:              GetStringConfiguration("DB_PASSWORD"),
        DbHost:                  GetStringConfiguration("DB_HOST"),
        DbPort:                  GetStringConfiguration("DB_PORT"),
        DbName:                  GetStringConfiguration("DB_NAME"),
        DbDriver:                GetStringConfiguration("DB_DRIVER"),
        DbMaxConnection:         GetIntConfiguration("DB_MAX_CONNECTION"),
        DbMaxIdleConnection:     GetIntConfiguration("DB_MAX_IDLE_CONNECTION"),
        DbMaxLifeTimeConnection: GetIntConfiguration("DB_MAX_LIFE_TIME_CONNECTION"),
    }
}
