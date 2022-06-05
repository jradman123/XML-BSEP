package main

import (
	_ "go.mongodb.org/mongo-driver/mongo"
	"post/module/startup"
	cfg "post/module/startup/config"
)

const (
	DATABASE   = "posts_service"
	COLLECTION = "postsData"
)

func main() {
	config := cfg.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
