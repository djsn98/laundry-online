package application

import (
	"LaundryOnlineAPI/config"
	"LaundryOnlineAPI/router"
	"fmt"

	g "github.com/incubus8/go/pkg/gin"
)

func init() {
	config.Load()
}

func StartApp() {
	var env = config.Config
	fmt.Println(env)

	address := env.HOST + ":" + env.PORT
	serverConfig := g.Config{
		ListenAddr: address,
		Handler:    router.New(),
		OnStarting: func() {
			fmt.Println("Server running on port " + env.PORT)
		},
	}
	g.Run(serverConfig)
}
