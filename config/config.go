package config

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

var configJSONFileName = "./config.json"

func init() {
	viper.SetConfigFile(configJSONFileName)
	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Connont find config file, %s", err)
	}
}

// NewConfig ...
func NewConfig() (defConfig *Config, err error) {
	defConfig = &Config{}
	appEnv := viper.GetString(`APP_ENV`)
	appPort := viper.GetInt(`APP_PORT`)
	debug := viper.GetBool(`APP_DEBUG`)

	dbHost := viper.GetString(`DB_HOST`)
	dbPort := viper.GetInt(`DB_PORT`)
	dbUser := viper.GetString(`DB_USER`)
	dbPassword := viper.GetString(`DB_PASS`)
	dbName := viper.GetString(`DB_NAME`)
	driverName := viper.GetString(`DB_DRIVER_NAME`)

	if appEnv == "" || appPort == 0 {
		err = fmt.Errorf("[CONFIG][Critical] Please check section APP on %s", configJSONFileName)
		return
	}

	defConfig.AppEnv = appEnv
	defConfig.AppPort = appPort
	defConfig.Debug = debug

	if dbHost == "" || dbPort == 0 || dbUser == "" || dbName == "" || driverName == "" {
		err = fmt.Errorf("[CONFIG][Critical] Please check section DB on %s", configJSONFileName)
		return
	}

	dbConfig := &DB{
		Host:       dbHost,
		Port:       dbPort,
		Username:   dbUser,
		Password:   dbPassword,
		Name:       dbName,
		DriverName: driverName,
	}

	defConfig.DB = dbConfig

	return defConfig, nil
}
