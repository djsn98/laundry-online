package application

import (
	"LaundryOnlineAPI/config"
	"LaundryOnlineAPI/router"
	"fmt"

	g "github.com/incubus8/go/pkg/gin"
)

func init() {
	// Load data from file "env" to var Config
	// Var Config there is in "./config/config.go"
	config.Load()
}

func StartApp() {
	// Get Config
	var env = config.Config
	fmt.Println(env)

	// Setup server
	address := env.HOST + ":" + env.PORT
	serverConfig := g.Config{
		ListenAddr: address,
		Handler:    router.New(),
		OnStarting: func() {
			fmt.Println("Server running on port " + env.PORT)
		},
	}

	// Run server
	g.Run(serverConfig)
}
