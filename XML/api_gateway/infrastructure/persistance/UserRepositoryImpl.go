package persistance

import (
	"context"
	"errors"
	"gateway/module/domain/model"
	"gateway/module/domain/repositories"
	"gorm.io/gorm"
	tracer "monitoring/module"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepositoryImpl(db *gorm.DB) repositories.UserRepository {
	return &UserRepositoryImpl{db: db}
}

func (r UserRepositoryImpl) GetUsers() ([]model.User, error) {
	var users []model.User
	r.db.Select("*").Find(&users)
	return users, nil
}

func (r UserRepositoryImpl) GetByUsername(ctx context.Context, username string) (*model.User, error) {
	span := tracer.StartSpanFromContext(ctx, "getByUsernameRepository")
	defer span.Finish()
	user := &model.User{}
	if r.db.First(&user, "username = ?", username).RowsAffected == 0 {
		return nil, errors.New("user not found")

	}
	return user, nil
}

func (r UserRepositoryImpl) UserExists(username string) error {
	if r.db.First(&model.User{}, "username = ?", username).RowsAffected == 0 {
		return errors.New("user does not exist")

	}
	return nil
}

func (r UserRepositoryImpl) GetUserSalt(username string) (string, error) {
	var result string = ""
	r.db.Table("users").Select("salt").Where("username = ?", username).Scan(&result)
	if result == "" {
		return "", errors.New("User salt not found!")
	}
	return result, nil
}

func (r UserRepositoryImpl) GetUserRole(username string, ctx context.Context) (string, error) {
	span := tracer.StartSpanFromContext(ctx, "getUserRoleRepository")
	defer span.Finish()

	var result int
	r.db.Table("users").Select("role").Where("username = ?", username).Scan(&result)

	if result == 0 {
		return "Regular", nil
	}
	if result == 1 {
		return "Admin", nil
	}
	if result == 2 {
		return "Agent", nil
	}
	return "", errors.New("User role not found for username" + username)
}
