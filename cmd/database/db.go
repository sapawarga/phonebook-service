package database

import (
	"fmt"
	"log"
	"net/url"

	"github.com/jmoiron/sqlx"
	"github.com/sapawarga/phonebook-service/config"
)

// NewConnection ...
func NewConnection(config *config.DB) *sqlx.DB {
	var err error
	connection := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", config.Username, config.Password, config.Host, config.Port, config.Name)
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Jakarta")
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())

	db, err := sqlx.Connect(config.DriverName, dsn)
	if err != nil {
		log.Panic("[CONFIG DB] error connect :", err)
	}

	if err = db.Ping(); err != nil {
		log.Panic("[CONFIG DB] error ping :", err)
	}

	return db
}
