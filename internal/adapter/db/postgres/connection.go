package postgres

import (
	"fmt"
	"log"
	"time"

	"github.com/rogalni/cng-hello-backend/internal/adapter/db/postgres/model"
	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	maxCon = 100
)

var (
	DBConn *gorm.DB
)

func InitConnection(host string, user string, password string, dbName string, port string) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbName, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt:     true,
		QueryFields:     true,
		CreateBatchSize: 50,
		Logger:          logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatalf("Could not connect to db: %v\n", err)
	}

	if err := db.Use(otelgorm.NewPlugin()); err != nil {
		log.Fatalf("Could not use plugin otelgorm: %v\n", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Could not connect to sql db: %v", err)
	}

	sqlDB.SetMaxOpenConns(maxCon)
	sqlDB.SetMaxIdleConns(int(float64(maxCon) * 0.1))
	sqlDB.SetConnMaxLifetime(time.Hour)
	// Ensure that the connection is established
	// Disabled in favour of health endpoints
	// if err := sqlDB.Ping(); err != nil {
	// 	panic(err)
	// }

	db.AutoMigrate(&model.Message{})
	return db
}
