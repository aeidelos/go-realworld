package configs

type DbConfig struct {
    DbUser     string
    DbPassword string
    DbHost     string
    DbPort     string
}

func NewDbConfig() DbConfig {
    return DbConfig{
        DbUser:     GetStringConfiguration("DB_USER"),
        DbPassword: GetStringConfiguration("DB_PASSWORD"),
        DbHost:     GetStringConfiguration("DB_HOST"),
        DbPort:     GetStringConfiguration("DB_PORT"),
    }
}
