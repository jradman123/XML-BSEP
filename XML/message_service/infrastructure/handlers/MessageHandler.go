package handlers

import (
	"common/module/interceptor"
	"common/module/logger"
	pb "common/module/proto/message_service"
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"message/module/application"
	"message/module/infrastructure/api"
)

type MessageHandler struct {
	messageService *application.MessageService
	userService    *application.UserService
	logInfo        *logger.Logger
	logError       *logger.Logger
}

func NewMessageHandler(messageService *application.MessageService, userService *application.UserService, logInfo *logger.Logger, logError *logger.Logger) *MessageHandler {
	return &MessageHandler{messageService: messageService, userService: userService, logInfo: logInfo, logError: logError}
}

func (m MessageHandler) MustEmbedUnimplementedMessageServiceServer() {
}

func (m MessageHandler) GetAllSent(ctx context.Context, request *pb.GetRequest) (*pb.GetMultipleResponse, error) {
	sender, err := m.userService.GetByUsername(request.Username)
	if err != nil {
		m.logError.Logger.WithFields(logrus.Fields{
			"userId": request.Username,
		}).Errorf("No user in database")
		return nil, err
	}
	messages, err := m.messageService.GetAllSent(sender[0].UserId.String())
	response := &pb.GetMultipleResponse{Messages: []*pb.Message{}}
	for _, message := range messages {
		receiver, _ := m.userService.GetByUsername(message.ReceiverId)
		current := api.MapMessageReply(message, receiver[0].UserId.String(), sender[0].UserId.String())
		response.Messages = append(response.Messages, current)
	}

	return response, nil
}

func (m MessageHandler) GetAllReceived(ctx context.Context, request *pb.GetRequest) (*pb.GetMultipleResponse, error) {
	receiver, err := m.userService.GetByUsername(request.Username)
	if err != nil {
		m.logError.Logger.WithFields(logrus.Fields{
			"userId": request.Username,
		}).Errorf("No user in database")
		return nil, err
	}
	messages, err := m.messageService.GetAllReceived(receiver[0].UserId.String())
	response := &pb.GetMultipleResponse{Messages: []*pb.Message{}}
	for _, message := range messages {
		sender, _ := m.userService.GetByUsername(message.SenderId)
		current := api.MapMessageReply(message, receiver[0].UserId.String(), sender[0].UserId.String())
		response.Messages = append(response.Messages, current)
	}

	return response, nil
}

func (m MessageHandler) SendMessage(ctx context.Context, request *pb.SendMessageRequest) (*pb.Empty, error) {
	userNameCtx := fmt.Sprintf(ctx.Value(interceptor.LoggedInUserKey{}).(string))
	userSender, _ := m.userService.GetByUsername(request.Message.SenderUsername)
	userReceiver, _ := m.userService.GetByUsername(request.Message.ReceiverUsername)

	message := api.MapNewMessage(request.Message.MessageText, userSender[0].UserId.String(), userReceiver[0].UserId.String())
	err := m.messageService.SendMessage(message)
	if err != nil {
		m.logError.Logger.WithFields(logrus.Fields{
			"userId": userNameCtx,
		}).Errorf("Can not send message")
		return nil, err
	}

	//TODO: Vjv treba vratiti tako da bih update listu poslatih
	return &pb.Empty{}, nil
}
