package main

import (
	startup "gateway/module/startup"
	config "gateway/module/startup/config"
)

func main() {
	config := config.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
