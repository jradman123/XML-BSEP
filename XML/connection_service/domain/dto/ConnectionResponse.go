package dto

type ConnectionResponse struct {
	UserOneUID       string //sender
	UserTwoUID       string //receiver
	ConnectionStatus string
}

