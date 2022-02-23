package config

import (
	"os"

	"github.com/joho/godotenv"
)

func Load() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err.Error())
	}

	Config.HOST = os.Getenv("HOST")
	Config.PORT = os.Getenv("PORT")
	Config.DBHOST = os.Getenv("DBHOST")
	Config.DBPORT = os.Getenv("DBPORT")
	Config.DBNAME = os.Getenv("DBNAME")
	Config.DBUSER = os.Getenv("DBUSER")
	Config.DBPASS = os.Getenv("DBPASS")
}
