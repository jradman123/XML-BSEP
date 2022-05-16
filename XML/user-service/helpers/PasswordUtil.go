package helpers

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type PasswordUtil struct {
}

func (util *PasswordUtil) IsValidPassword(userInput string) bool {
	uppercase := `[A-Z]{1}`
	lowercase := `[a-z]{1}`
	number := `[0-9]{1}`
	specialCharacters := `[!"#$@%&()*<>+\_|~]{1}`
	whiteSpace := ` {1}`

	if matched, err := regexp.MatchString(uppercase, userInput); !matched || err != nil {
		fmt.Println("Your password should contain at least one uppercase letter.")
		return false
	}

	if matched, err := regexp.MatchString(lowercase, userInput); !matched || err != nil {
		fmt.Println("Your password should contain at least one lowercase letter.")
		return false
	}

	if matched, err := regexp.MatchString(number, userInput); !matched || err != nil {
		fmt.Println("Your password should contain at least one number.")
		return false
	}

	if matched, err := regexp.MatchString(specialCharacters, userInput); !matched || err != nil {
		fmt.Println("Your password should contain at least one special character.")
		return false
	}

	if matched, err := regexp.MatchString(whiteSpace, userInput); matched || err != nil {
		fmt.Println("Your password shouldn't contain spaces.")
		return false
	}
	fmt.Println("Thanks, you have entered a password in a valid format!")
	return true
}

func (util *PasswordUtil) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(bytes), err
}

func (util *PasswordUtil) GeneratePasswordWithSalt(plainPassword string) (string, string) {
	var sb strings.Builder
	salt := uuid.New().String()
	sb.WriteString(plainPassword)
	sb.WriteString(salt)
	passwordWithSalt := sb.String()
	hashedPassword, _ := util.HashPassword(passwordWithSalt)
	return salt, hashedPassword
}

func (util *PasswordUtil) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
