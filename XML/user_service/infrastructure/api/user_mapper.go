package api

import (
	pb "common/module/proto/user_service"
	"fmt"
	"github.com/google/uuid"
	"time"
	"user/module/domain/dto"
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
		IsConfirmed:   user.IsConfirmed,
	}
	return usersPb
}

func MapUserToPbResponseUser(user *model.User) *pb.RegisteredUser {
	usersPb := &pb.RegisteredUser{
		Username:  user.Username,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}
	return usersPb
}
func MapDtoToUser(userPb *dto.NewUser) *model.User {
	userD := &model.User{
		ID:            uuid.New(),
		FirstName:     userPb.FirstName,
		LastName:      userPb.LastName,
		Username:      userPb.Username,
		Email:         userPb.Email,
		PhoneNumber:   userPb.PhoneNumber,
		Gender:        mapGenderToModel(userPb.Gender),
		DateOfBirth:   mapToDate(userPb.DateOfBirth),
		Password:      userPb.Password,
		Role:          model.Regular,
		IsConfirmed:   false,
		RecoveryEmail: userPb.RecoveryEmail,
	}
	return userD
}

func MapPbUserToNewUserDto(userPb *pb.RegisterUserRequest) *dto.NewUser {
	fmt.Printf("Eo ga userPb: %v", userPb)
	userD := &dto.NewUser{
		FirstName:     userPb.UserRequest.FirstName,
		LastName:      userPb.UserRequest.LastName,
		Username:      userPb.UserRequest.Username,
		Email:         userPb.UserRequest.Email,
		PhoneNumber:   userPb.UserRequest.PhoneNumber,
		Gender:        userPb.UserRequest.Gender,
		DateOfBirth:   userPb.UserRequest.DateOfBirth,
		Password:      userPb.UserRequest.Password,
		RecoveryEmail: userPb.UserRequest.RecoveryEmail,
	}
	return userD
}

func mapToDate(birth string) time.Time {
	layout := "2006-01-02T15:04:05.000Z"
	dateOfBirth, _ := time.Parse(layout, birth)
	return dateOfBirth

}
func mapGenderToModel(gender string) model.Gender {
	switch gender {
	case "MALE":
		return model.MALE
	case "FEMALE":
		return model.FEMALE
	case "OTHER":
		return model.OTHER
	}
	return model.OTHER
}

func MapPbToUserActivateRequest(request *pb.ActivationRequest) *dto.UserActivateRequest {
	requestDTO := &dto.UserActivateRequest{
		Code:     request.Account.Code,
		Username: request.Account.Username,
	}
	return requestDTO
}
