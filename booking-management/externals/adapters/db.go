package adapters

import (
	"booking-management/utils/config"
	"database/sql"
	"strconv"
	"time"

	"github.com/go-sql-driver/mysql"
)

func NewDB(conf config.Database) (*sql.DB, error) {

	config := mysql.Config{
		User:                 conf.User,
		Passwd:               conf.Password,
		Net:                  "tcp",
		Addr:                 conf.Host + ":" + strconv.Itoa(conf.Port),
		DBName:               conf.Database,
		AllowNativePasswords: true,
		ParseTime:            true,
		Timeout:              time.Duration(60000) * time.Millisecond,
		ReadTimeout:          time.Duration(60000) * time.Millisecond,
		WriteTimeout:         time.Duration(60000) * time.Millisecond,
	}

	conn := config.FormatDSN()

	db, err := sql.Open("mysql", conn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(time.Duration(600000) * time.Millisecond)
	db.SetConnMaxIdleTime(time.Duration(600000) * time.Millisecond)

	return db, nil
}
