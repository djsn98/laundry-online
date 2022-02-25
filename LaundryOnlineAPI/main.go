package main

import (
	"LaundryOnlineAPI/application"
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
	application.StartApp()
}
