package config

// DB ...
type DB struct {
	Host       string `env:"DB_HOST,required"`
	Port       int    `env:"DB_PORT,required"`
	Username   string `env:"DB_USER,required"`
	Password   string `env:"DB_PASS"`
	DriverName string `env:"DB_DRIVER_NAME,required"`
	Name       string `env:"DB_NAME,required"`
}

// Config ...
type Config struct {
	AppPort int    `env:"APP_PORT,required"`
	AppEnv  string `env:"APP_ENV,required"`
	Debug   bool   `env:"APP_DEBUG,required"`
	DB      *DB
}
