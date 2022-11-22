package persistance

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	tracer "monitoring/module"
	"user/module/domain/model"
	"user/module/domain/repositories"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func (r UserRepositoryImpl) ActivateUserAccount(user *model.User, ctx context.Context) (bool, error) {
	span := tracer.StartSpanFromContext(ctx, "activateUserAccountRepo")
	defer span.Finish()

	result := r.db.Model(&user).Update("is_confirmed", true)
	fmt.Print(result)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

func (r UserRepositoryImpl) EditUserDetails(user *model.User, ctx context.Context) (bool, error) {
	span := tracer.StartSpanFromContext(ctx, "editUserDetailsRepository")
	defer span.Finish()

	result := r.db.Model(&user).Updates(&user)
	r.db.Model(&user).Association("Skills").Replace(user.Skills)
	r.db.Model(&user).Association("Interests").Replace(user.Interests)
	r.db.Model(&user).Association("Educations").Replace(user.Educations)
	r.db.Model(&user).Association("Experiences").Replace(user.Experiences)
	r.db.Where("user_id is null").Delete(&model.Skill{})
	r.db.Where("user_id is null").Delete(&model.Interest{})
	r.db.Where("user_id is null").Delete(&model.Education{})
	r.db.Where("user_id is null").Delete(&model.Experience{})
	fmt.Print(result)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

func NewUserRepositoryImpl(db *gorm.DB) repositories.UserRepository {
	return &UserRepositoryImpl{db: db}
}

func (r UserRepositoryImpl) GetUsers(ctx context.Context) ([]model.User, error) {
	span := tracer.StartSpanFromContext(ctx, "getUsersRepository")
	defer span.Finish()

	var users []model.User
	r.db.Select("*").Find(&users)
	return users, nil
}

func (r UserRepositoryImpl) GetByUsername(ctx context.Context, username string) (*model.User, error) {
	span := tracer.StartSpanFromContext(ctx, "getByUsernameRepository")
	defer span.Finish()

	user := &model.User{}
	if r.db.Preload("Skills").Preload("Interests").Preload("Educations").Preload("Experiences").First(&user, "username = ?", username).RowsAffected == 0 {
		return nil, errors.New("user not found")

	}
	return user, nil
}

func (r UserRepositoryImpl) CreateRegisteredUser(user *model.User, ctx context.Context) (*model.User, error) {
	span := tracer.StartSpanFromContext(ctx, "createRegisteredUserRepository")
	defer span.Finish()

	result := r.db.Create(&user)
	fmt.Print(result)
	regUser := &model.User{}
	r.db.First(&regUser, user.ID)
	return regUser, result.Error
}

func (r UserRepositoryImpl) UserExists(username string, ctx context.Context) error {
	span := tracer.StartSpanFromContext(ctx, "userExistsRepository")
	defer span.Finish()
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

func (r UserRepositoryImpl) ChangePassword(user *model.User, password string, ctx context.Context) error {
	span := tracer.StartSpanFromContext(ctx, "changePasswordRepository")
	defer span.Finish()

	result := r.db.Model(&user).Update("password", password)
	fmt.Print(result)
	return result.Error
}

func (r UserRepositoryImpl) ChangeProfileStatus(user *model.User) (bool, error) {
	result := r.db.Model(&user).Updates(&user)
	fmt.Print(result)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

func (r UserRepositoryImpl) GetById(ctx context.Context, id uuid.UUID) (*model.User, error) {
	span := tracer.StartSpanFromContext(ctx, "getByIdRepository")
	defer span.Finish()

	user := &model.User{}
	if r.db.Preload("Skills").Preload("Interests").Preload("Educations").Preload("Experiences").First(&user, "id = ?", id).RowsAffected == 0 {
		return nil, errors.New("user not found")

	}
	return user, nil
}

func (r UserRepositoryImpl) UpdateEmail(ctx context.Context, user *model.User) (bool, error) {
	span := tracer.StartSpanFromContext(ctx, "updateEmailRepository")
	defer span.Finish()

	result := r.db.Model(&user).Update("email", user.Email)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

func (r UserRepositoryImpl) UpdateUsername(ctx context.Context, user *model.User) (bool, error) {
	span := tracer.StartSpanFromContext(ctx, "updateUsernameRepository")
	defer span.Finish()
	result := r.db.Model(&user).Update("username", user.Username)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}
