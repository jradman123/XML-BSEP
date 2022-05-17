package main

import (
	config "api-gateway/module/startup"
	startup "api-gateway/module/startup/config"
)

func main() {
	config := config.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
