package service

import (
	"time"
	"user/module/model"
	"user/module/repository"

	"github.com/google/uuid"
)

type RegisteredUserService struct {
	Repo *repository.RegisteredUserRepository
}

func (service *RegisteredUserService) CreateRegisteredUser(username string, password string, email string, phone string, firstName string, lastName string, gender model.Gender, role model.UserType, dateOfBirth time.Time, question string, answer string) (string, error) {
	user := model.User{
		ID:           uuid.New(),
		Username:     username,
		Password:     password,
		Email:        email,
		PhoneNumber:  phone,
		FirstName:    firstName,
		LastName:     lastName,
		Gender:       gender,
		Role:         role,
		IsConfirmed:  false,
		DateOfBirth:  dateOfBirth,
		Question:     question,
		HashedAnswer: answer,
	}
	mail, err := service.Repo.CreateRegisteredUser(&user)
	if err != nil {
		return mail, err
	}

	//TODO: send confirmation mail
	//koriste redis bazu gdje privremeno cuvaju zahteve za registraciju
	//a ovo sto serijalizuju mzd kasnije tek upisuju u bazu
	//expiration  := 1000000000 * 3600 * 2 //2h
	//serializedUser, err := serialize(user)
	// err = s.RedisUsecase.AddKeyValueSet(context, redisKey, serializedUser, time.Duration(expiration));
	// if err != nil {
	// 	return err
	// }
	// confirmationCode := helpers.RandomStringGenerator(8)
	// hashedConfirmationCode, err := helpers.Hash(confirmationCode)
	// user.hashedConfirmationCode = hashedConfirmationCode
	// if err != nil {
	// 	s.logger.Logger.Errorf("error while registering user, error %v\n", err)
	// 	return err
	// }
	return mail, nil
}
