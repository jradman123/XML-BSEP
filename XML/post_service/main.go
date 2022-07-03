package main

import (
	_ "go.mongodb.org/mongo-driver/mongo"
	"post/module/startup"
	cfg "post/module/startup/config"
)

func main() {
	config := cfg.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
