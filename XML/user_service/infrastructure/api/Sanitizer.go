package api

import (
	common "common/module"
	"github.com/microcosm-cc/bluemonday"

	"strings"
	"user/module/domain/dto"
)

func SanitizeUser(newUser *dto.NewUser) (error string) {
	policy := bluemonday.UGCPolicy()
	newUser.Username = strings.TrimSpace(policy.Sanitize(newUser.Username))
	newUser.FirstName = strings.TrimSpace(policy.Sanitize(newUser.FirstName))
	newUser.LastName = strings.TrimSpace(policy.Sanitize(newUser.LastName))
	newUser.Email = strings.TrimSpace(policy.Sanitize(newUser.Email))
	newUser.Password = strings.TrimSpace(policy.Sanitize(newUser.Password))
	newUser.Gender = strings.TrimSpace(policy.Sanitize(newUser.Gender))
	newUser.DateOfBirth = strings.TrimSpace(policy.Sanitize(newUser.DateOfBirth))
	newUser.PhoneNumber = strings.TrimSpace(policy.Sanitize(newUser.PhoneNumber))
	newUser.RecoveryEmail = strings.TrimSpace(policy.Sanitize(newUser.RecoveryEmail))

	p1 := common.BadUsername(newUser.Username)
	p2 := common.BadName(newUser.FirstName)
	p3 := common.BadName(newUser.LastName)
	p4 := common.BadEmail(newUser.Email)
	p5 := common.BadText(newUser.Gender)
	p6 := common.BadDate(newUser.DateOfBirth)
	p7 := common.BadNumber(newUser.PhoneNumber)
	p8 := common.BadEmail(newUser.RecoveryEmail)
	p9 := common.BadPassword(newUser.Password)

	if newUser.Username == "" || newUser.FirstName == "" || newUser.LastName == "" ||
		newUser.Gender == "" || newUser.DateOfBirth == "" || newUser.PhoneNumber == "" ||
		newUser.Password == "" || newUser.Email == "" || newUser.RecoveryEmail == "" {
		return "XSS"

	} else if p1 || p2 || p3 || p4 || p5 || p6 || p7 || p8 || p9 {
		return "POSSIBLE INJECTION"

	}
	return ""

}
