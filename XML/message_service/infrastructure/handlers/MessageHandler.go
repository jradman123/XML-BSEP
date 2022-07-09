package handlers

import (
	"common/module/logger"
	pb "common/module/proto/message_service"
	"context"
	"fmt"
	"github.com/pusher/pusher-http-go"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"message/module/application"
	"message/module/domain/model"
	"message/module/infrastructure/api"
	"time"
)

type MessageHandler struct {
	messageService      *application.MessageService
	userService         *application.UserService
	notificationService *application.NotificationService
	logInfo             *logger.Logger
	logError            *logger.Logger
	pusher              *pusher.Client
}

func NewMessageHandler(messageService *application.MessageService, userService *application.UserService, notificationService *application.NotificationService, logInfo *logger.Logger, logError *logger.Logger, pusher *pusher.Client) *MessageHandler {
	return &MessageHandler{messageService: messageService, userService: userService, notificationService: notificationService, logInfo: logInfo, logError: logError, pusher: pusher}
}

func (m MessageHandler) MustEmbedUnimplementedMessageServiceServer() {
}

func (m MessageHandler) GetAllSent(_ context.Context, request *pb.GetRequest) (*pb.GetMultipleResponse, error) {
	fmt.Println("usao u hendler get all sent")
	fmt.Println("dobio username " + request.Username)
	sender, err := m.userService.GetByUsername(request.Username)
	if err != nil {
		m.logError.Logger.WithFields(logrus.Fields{
			"userId": request.Username,
		}).Errorf("No user in database")
		fmt.Println("nemas ovog usera u bAzi")
		return nil, err
	}
	fmt.Println("sender[0].UserId")
	fmt.Println(sender[0].UserId)
	messages, err := m.messageService.GetAllSent(sender[0].UserId)

	fmt.Println(messages)

	fmt.Println(messages)
	response := &pb.GetMultipleResponse{Messages: []*pb.Message{}}
	for _, message := range messages {
		receiver, _ := m.userService.GetById(message.ReceiverId)
		current := api.MapMessageReply(message, receiver[0].Username, sender[0].Username)
		response.Messages = append(response.Messages, current)
	}

	return response, nil
}

func (m MessageHandler) GetAllReceived(_ context.Context, request *pb.GetRequest) (*pb.GetMultipleResponse, error) {
	fmt.Println("usao u hendler get all receiver")
	fmt.Println("dobio username " + request.Username)

	receiver, err := m.userService.GetByUsername(request.Username)
	if err != nil {
		m.logError.Logger.WithFields(logrus.Fields{
			"userId": request.Username,
		}).Errorf("No user in database")
		fmt.Println("nemas ovog usera u bAzi")
		return nil, err
	}
	messages, err := m.messageService.GetAllReceived(receiver[0].UserId)
	response := &pb.GetMultipleResponse{Messages: []*pb.Message{}}
	for _, message := range messages {
		sender, _ := m.userService.GetById(message.SenderId)
		current := api.MapMessageReply(message, receiver[0].Username, sender[0].Username)
		response.Messages = append(response.Messages, current)
	}

	return response, nil
}

func (m MessageHandler) SendMessage(_ context.Context, request *pb.SendMessageRequest) (*pb.MessageSentResponse, error) {

	userSender, _ := m.userService.GetByUsername(request.Message.SenderUsername)
	userReceiver, _ := m.userService.GetByUsername(request.Message.ReceiverUsername)

	modell := api.MapNewMessage(request.Message.MessageText, userReceiver[0].UserId, userSender[0].UserId)
	message, err := m.messageService.SendMessage(modell)

	if err != nil {
		m.logError.Logger.WithFields(logrus.Fields{
			"userId": request.Message.SenderUsername,
		}).Errorf("Can not send message")
		return nil, err
	}

	nnn := &model.Notification{
		Id:               primitive.NewObjectID(),
		Timestamp:        time.Now(),
		Content:          request.Message.SenderUsername + " sent you a message.",
		NotificationFrom: request.Message.SenderUsername,
		NotificationTo:   request.Message.ReceiverUsername,
		Read:             false,
		RedirectPath:     "/chatbox",
		Type:             model.MESSAGE,
	}
	m.notificationService.Create(nnn)
	m.pusher.Trigger("messages", "message", request.Message)
	response := api.MapMessageReply(message, request.Message.ReceiverUsername, request.Message.SenderUsername)
	return &pb.MessageSentResponse{Message: response}, nil
}
