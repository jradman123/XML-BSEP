package model

type UserType int

const (
	ADMIN UserType = iota
	REGISTERED_USER
	AGENT
)
