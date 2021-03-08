package config

// DB ...
type DB struct {
	Host       string
	Port       int
	Username   string
	Password   string
	DriverName string
	Name       string
}

// Config ...
type Config struct {
	AppPort string
	AppEnv  string
	Debug   bool
	DB      *DB
}
