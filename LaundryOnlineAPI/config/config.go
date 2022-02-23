package config

type ConfigModel struct {
	HOST   string
	PORT   string
	DBNAME string
	DBUSER string
	DBHOST string
	DBPORT string
	DBPASS string
}

var Config ConfigModel
