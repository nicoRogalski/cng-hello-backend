package postgres

import (
	"fmt"
	"time"

	"github.com/rogalni/cng-hello-backend/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DBConn *gorm.DB
)

func InitConnection() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.App.PostgresHost, config.App.PostgresUser, config.App.PostresPassword, config.App.PostgresDb, config.App.PostgresPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt: true,
	})
	if err != nil {
		panic("Could not connect to db")
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic("Could not connect to sql db")
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	// Ensure that the connection is established
	e := sqlDB.Ping()
	if e != nil {
		panic(e)
	}
	DBConn = db
}
