package main

import (
	"user/module/startup"
	cf "user/module/startup/config"
)

func main() {
	config := cf.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
