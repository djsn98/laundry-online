package database

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	// var env = config.Config
	// fmt.Print(env)
	// dbConf := fmt.Sprintf(
	// 	"host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
	// 	env.DBHOST,
	// 	env.DBPORT,
	// 	env.DBUSER,
	// 	env.DBNAME,
	// 	env.DBPASS,
	// )
	sqlDB, err := sql.Open("postgres", "host=localhost port=5432 user=dennis dbname=laundry_online_api_db sslmode=disable password=susah13234")
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
