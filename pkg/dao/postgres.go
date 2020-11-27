package dao

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	// PgConn it is safe for goroutines. Use PgConn for all postgres request.
	PgConn = &gorm.DB{}
)

func ConnectPG(host, port, user, dbname, passwd string) (err error) {
	dbEndPoint := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable connect_timeout=2", host, port, user, passwd, dbname)
	PgConn, err = gorm.Open(postgres.Open(dbEndPoint), &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return
	}

	db, err := PgConn.DB()
	if err != nil {
		return
	}
	db.SetConnMaxLifetime(time.Minute * 5)
	// db.SetMaxIdleConns(20)
	// db.SetMaxOpenConns(500)
	return
}
