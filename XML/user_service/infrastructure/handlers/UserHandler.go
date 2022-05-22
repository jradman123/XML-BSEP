package handlers

import (
	pb "common/module/proto/user_service"
	"context"
	"gopkg.in/go-playground/validator.v9"
	"log"
	"user/module/application/helpers"
	"user/module/application/services"
	"user/module/infrastructure/api"
)

type UserHandler struct {
	l            *log.Logger
	service      *services.UserService
	jsonConv     *helpers.JsonConverters
	validator    *validator.Validate
	passwordUtil *helpers.PasswordUtil
}

func NewUserHandler(l *log.Logger, service *services.UserService, jsonConv *helpers.JsonConverters, validator *validator.Validate,
	passwordUtil *helpers.PasswordUtil) *UserHandler {
	return &UserHandler{l, service, jsonConv, validator, passwordUtil}
}

func (u UserHandler) MustEmbedUnimplementedUserServiceServer() {
	u.l.Println("Handling MustEmbedUnimplementedUserServiceServer Users")
}

func (u UserHandler) GetAll(ctx context.Context, request *pb.EmptyRequest) (*pb.GetAllResponse, error) {
	u.l.Println("Handling GetAll Users")
	users, err := u.service.GetUsers()
	if err != nil {
		return nil, err
	}
	response := &pb.GetAllResponse{
		Users: []*pb.User{},
	}
	for _, user := range users {
		current := api.MapProduct(&user)
		response.Users = append(response.Users, current)
	}
	return response, nil
}

func (u UserHandler) UpdateUser(ctx context.Context, request *pb.UpdateRequest) (*pb.UpdateUserResponse, error) {
	u.l.Println("Handling UpdateUser Users")
	return &pb.UpdateUserResponse{UpdatedUser: nil}, nil
}

func (u UserHandler) RegisterUser(ctx context.Context, request *pb.RegisterUserRequest) (*pb.RegisterResponse, error) {
	u.l.Println("Handling RegisterUser Users")
	return &pb.RegisterResponse{User: nil}, nil
}

func (u UserHandler) Login(ctx context.Context, request *pb.LoginUserRequest) (*pb.LoginResponse, error) {
	u.l.Println("Handling Login Users")
	return &pb.LoginResponse{Token: ""}, nil
}
