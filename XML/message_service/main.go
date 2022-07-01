package main

import (
	_ "go.mongodb.org/mongo-driver/mongo"
	"message/module/startup"
	cfg "message/module/startup/config"
)

const (
	DATABASE   = "message_service"
	COLLECTION = "messagesData"
)

func main() {
	config := cfg.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
