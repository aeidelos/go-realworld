package configs

type Db struct {
    DbUser     string
    DbPassword string
    DbHost     string
    DbPort     string
    DbName     string
}

func NewDbConfig() Db {
    return Db{
        DbUser:     GetStringConfiguration("DB_USER"),
        DbPassword: GetStringConfiguration("DB_PASSWORD"),
        DbHost:     GetStringConfiguration("DB_HOST"),
        DbPort:     GetStringConfiguration("DB_PORT"),
        DbName:     GetStringConfiguration("DB_NAME"),
    }
}
