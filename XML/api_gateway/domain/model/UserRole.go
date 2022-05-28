package model

type Role int

const (
	Regular Role = iota
	Admin
	Agent
)
