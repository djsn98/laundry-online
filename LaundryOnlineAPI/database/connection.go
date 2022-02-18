package database

import (
	"database/sql"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	sqlDB, err := sql.Open("postgres", "host=localhost user=postgres password=susah13234 dbname=learn_grpc port=5432 sslmode=disable TimeZone=Asia/Jakarta")
	if err != nil {
		panic(err)
	}

	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetMaxOpenConns(20)
	sqlDB.SetConnMaxLifetime(60 * time.Minute)
	sqlDB.SetConnMaxIdleTime(10 * time.Minute)

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return gormDB
}

var Connection *gorm.DB

func init() {
	Connection = Connect()
}
