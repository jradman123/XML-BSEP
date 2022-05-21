package helpers

import (
	"fmt"
	"regexp"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	enTranslations "github.com/go-playground/validator/translations/en"
	"gopkg.in/go-playground/validator.v9"
)

type customValidator struct {
	Validator *validator.Validate
}

const (
	FIRST_NAME = "^[A-Z]{1}[a-z]+$"
	SURNAME    = "^[A-Z]{1}[a-z]+$"
	USERNAME   = "^[a-z]{4,}[0-9]*$"
	EMAIL      = "([a-zA-Z0-9]+)@([a-zA-Z0-9\\.]+)\\.([a-zA-Z0-9]+)"
	PHONE      = "^[0-9]{3,6}(\\/).[0-9]+$"
)

const (
	FIRST_NAME_ERROR_MSG = "{0} must be in valid format. First letter is uppercase"
	SURNAME_ERR_MSG      = "{0} must be in valid format. First letter is uppercase"
	EMAIL_ERR_MSG        = "{0} must be in valid format. Must contain @ and . Ex: mail@mail.com"
	USERNAME_ERR_MSG     = "{0} must be in valid format. At least 4 small letters with numbers"
	PHONE_ERR_MSG        = "{0} must be in valid format. Must contain 3-6 digits with '/' and more digits"
)

func NewCustomValidator() *customValidator {
	cv := &customValidator{validator.New()}
	err := registerNameValdation(cv)
	err = registerSurnameValidation(cv)
	err = registerEmailValidation(cv)
	err = registerUsernameValidation(cv)
	err = registerPhoneValidation(cv)
	if err != nil {
		return &customValidator{}
	}

	return cv
}

func (cv *customValidator) RegisterEnTranslation() (ut.Translator, error) {
	english := en.New()
	uni := ut.New(english, english)
	trans, _ := uni.GetTranslator("en")
	registerEnNameTranslation(trans, cv)
	registerEnSurnameTranslation(trans, cv)
	registerEnEmailTranslation(trans, cv)
	registerEnUsernameTranslation(trans, cv)
	registerEnPhoneTranslation(trans, cv)

	return trans, enTranslations.RegisterDefaultTranslations(cv.Validator, trans)
}

func (cv *customValidator) TranslateError(err error, translator ut.Translator) (translatedErrors []error) {
	if err == nil {
		return nil
	}
	validatorErrs := err.(validator.ValidationErrors)
	for _, e := range validatorErrs {
		translatedErr := fmt.Errorf(e.Translate(translator))
		translatedErrors = append(translatedErrors, translatedErr)
	}
	return translatedErrors
}

func (cv *customValidator) GetErrorsString(errs []error) (errsString []string) {
	for _, e := range errs {
		errsString = append(errsString, e.Error())
	}
	return errsString
}

func registerNameValdation(cv *customValidator) error {

	return cv.Validator.RegisterValidation("name", func(f1 validator.FieldLevel) bool {
		mathced, _ := regexp.Match(FIRST_NAME, []byte(f1.Field().String()))
		return mathced
	})

}

func registerSurnameValidation(cv *customValidator) error {
	return cv.Validator.RegisterValidation("surname", func(f1 validator.FieldLevel) bool {
		mathced, _ := regexp.Match(SURNAME, []byte(f1.Field().String()))
		return mathced
	})
}

func registerEnNameTranslation(tr ut.Translator, cv *customValidator) {
	_ = cv.Validator.RegisterTranslation("name", tr, func(ut ut.Translator) error {
		return ut.Add("name", FIRST_NAME_ERROR_MSG, true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("name", fe.Field())
		return t
	})
}

func registerEnSurnameTranslation(tr ut.Translator, cv *customValidator) {
	_ = cv.Validator.RegisterTranslation("surname", tr, func(ut ut.Translator) error {
		return ut.Add("surname", SURNAME_ERR_MSG, true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("surname", fe.Field())
		return t
	})
}

func registerEmailValidation(cv *customValidator) error {
	return cv.Validator.RegisterValidation("email", func(f1 validator.FieldLevel) bool {
		mathced, _ := regexp.Match(EMAIL, []byte(f1.Field().String()))
		return mathced
	})
}

func registerEnEmailTranslation(tr ut.Translator, cv *customValidator) {
	_ = cv.Validator.RegisterTranslation("email", tr, func(ut ut.Translator) error {
		return ut.Add("email", EMAIL_ERR_MSG, true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("email", fe.Field())
		return t
	})
}

func registerUsernameValidation(cv *customValidator) error {
	return cv.Validator.RegisterValidation("username", func(f1 validator.FieldLevel) bool {
		mathced, _ := regexp.Match(USERNAME, []byte(f1.Field().String()))
		return mathced
	})
}

func registerEnUsernameTranslation(tr ut.Translator, cv *customValidator) {
	_ = cv.Validator.RegisterTranslation("username", tr, func(ut ut.Translator) error {
		return ut.Add("username", USERNAME_ERR_MSG, true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("username", fe.Field())
		return t
	})
}

func registerPhoneValidation(cv *customValidator) error {
	return cv.Validator.RegisterValidation("phone", func(f1 validator.FieldLevel) bool {
		mathced, _ := regexp.Match(PHONE, []byte(f1.Field().String()))
		return mathced
	})
}

func registerEnPhoneTranslation(tr ut.Translator, cv *customValidator) {
	_ = cv.Validator.RegisterTranslation("phone", tr, func(ut ut.Translator) error {
		return ut.Add("phone", PHONE_ERR_MSG, true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("phone", fe.Field())
		return t
	})

}
