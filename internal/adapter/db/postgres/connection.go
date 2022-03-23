package postgres

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/rogalni/cng-hello-backend/config"
	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DBConn *gorm.DB
	SqlDb  *sql.DB
)

func InitConnection() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.App.PostgresHost, config.App.PostgresUser, config.App.PostresPassword, config.App.PostgresDb, config.App.PostgresPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt: true,
		Logger:      logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic("Could not connect to db")
	}

	if err := db.Use(otelgorm.NewPlugin()); err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic("Could not connect to sql db")
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	// Ensure that the connection is established
	// if err := sqlDB.Ping(); err != nil {
	// 	panic(err)
	// }
	DBConn = db
	SqlDb = sqlDB
}
