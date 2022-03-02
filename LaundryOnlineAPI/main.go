package main

import (
	"LaundryOnlineAPI/application" // Run init func to load Env config
	"LaundryOnlineAPI/database"
	"flag"
	"fmt"
)

func main() {
	// Run DB migration if migrate cli arg true
	var migrate = flag.Bool("migrate", false, "type migrate state")
	flag.Parse()

	if *migrate {
		database.Migrate()
		fmt.Println("DB migrate doned")
		return
	}
	// Run server app
	application.StartApp()
}
