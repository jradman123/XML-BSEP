package api

import (
	connectionPb "common/module/proto/connection_service"
	postPb "common/module/proto/posts_service"
	userPb "common/module/proto/user_service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func NewUserClient(serviceAddress string) userPb.UserServiceClient {
	conn, err := grpc.Dial(serviceAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to User service: %v", err)
	}
	return userPb.NewUserServiceClient(conn)
}

func NewPostClient(serviceAddress string) postPb.PostServiceClient {
	conn, err := grpc.Dial(serviceAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to Post service: %v", err)
	}
	return postPb.NewPostServiceClient(conn)
}

func NewConnectionClient(serviceAddress string) connectionPb.ConnectionServiceClient {
	conn, err := grpc.Dial(serviceAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to Connection service: %v", err)
	}
	return connectionPb.NewConnectionServiceClient(conn)
}
