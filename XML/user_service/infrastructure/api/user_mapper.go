package api

import (
	pb "common/module/proto/user_service"
	"user/module/domain/model"
)

func MapProduct(user *model.User) *pb.User {
	usersPb := &pb.User{
		Username:      user.Username,
		Password:      user.Password,
		Email:         user.Email,
		PhoneNumber:   user.PhoneNumber,
		FirstName:     user.FirstName,
		LastName:      user.LastName,
		Gender:        string(user.Gender),
		Role:          string(user.Role),
		DateOfBirth:   user.DateOfBirth.String(),
		RecoveryEmail: user.RecoveryEmail,
		IsConfirmed: user.IsConfirmed
	}
	return usersPb
}
