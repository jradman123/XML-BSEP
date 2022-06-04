package services

import (
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"log"
)

type ConnectionService struct {
	l         *log.Logger
	neoClient *neo4j.Driver
}

func NewConnectionService(l *log.Logger, neoClient *neo4j.Driver) *ConnectionService {
	return &ConnectionService{l, neoClient}
}
