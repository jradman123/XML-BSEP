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
		ProfileStatus: string(user.ProfileStatus),
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
	status := model.Public
	if userPb.ProfileStatus == "PRIVATE" {
		status = model.Private
	}

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
		ProfileStatus: status,
	}
	return userD
}

func MapPbUserDetailsToUser(userDetailsPb *pb.UserDetailsRequest) *dto.UserDetails {

	userD := &dto.UserDetails{
		FirstName:   userDetailsPb.UserDetails.FirstName,
		LastName:    userDetailsPb.UserDetails.LastName,
		Username:    userDetailsPb.UserDetails.Username,
		Email:       userDetailsPb.UserDetails.Email,
		PhoneNumber: userDetailsPb.UserDetails.PhoneNumber,
		Gender:      userDetailsPb.UserDetails.Gender,
		DateOfBirth: userDetailsPb.UserDetails.DateOfBirth,
		Biography:   userDetailsPb.UserDetails.Biography,
	}
	for i, s := range userDetailsPb.UserDetails.Educations {
		fmt.Println(i, s)
		ed := mapPbEducationToEducationDto(s)
		userD.Educations = append(userD.Educations, *ed)
	}
	for i, s := range userDetailsPb.UserDetails.Interests {
		fmt.Println(i, s)
		ed := mapPbEducationToInterestDto(s)
		userD.Interests = append(userD.Interests, *ed)
	}
	for i, s := range userDetailsPb.UserDetails.Skills {
		fmt.Println(i, s)
		ed := mapPbEducationToSkillDto(s)
		userD.Skills = append(userD.Skills, *ed)
	}
	for i, s := range userDetailsPb.UserDetails.Experiences {
		fmt.Println(i, s)
		ed := mapPbEducationToExperienceDto(s)
		userD.Experiences = append(userD.Experiences, *ed)
	}
	return userD
}
func mapPbEducationToEducationDto(e *pb.Education) *dto.EducationDto {
	education := &dto.EducationDto{
		Education: e.Education,
	}
	return education
}
func mapPbEducationToSkillDto(e *pb.Skill) *dto.SkillDto {
	skill := &dto.SkillDto{
		Skill: e.Skill,
	}
	return skill
}
func mapPbEducationToInterestDto(e *pb.Interest) *dto.InterestDto {
	interest := &dto.InterestDto{
		Interest: e.Interest,
	}
	return interest
}
func mapPbEducationToExperienceDto(e *pb.Experience) *dto.ExperienceDto {
	experience := &dto.ExperienceDto{
		Experience: e.Experience,
	}
	return experience
}

func MapUserDetailsDtoToUser(dto *dto.UserDetails, user *model.User) *model.User {
	user.Biography = dto.Biography
	user.FirstName = dto.FirstName
	user.LastName = dto.LastName
	user.Gender = mapGenderToModel(dto.Gender)
	user.PhoneNumber = dto.PhoneNumber
	user.DateOfBirth = mapToDate(dto.DateOfBirth)
	//KAKO SAD DA SVE NIZOVE PONISTIM DA BUDU PRAZNI NA POCETKU
	var skills []model.Skill
	var interests []model.Interest
	var educations []model.Education
	var experiences []model.Experience
	for i, s := range dto.Educations {
		fmt.Println(i, s)
		ed := mapEducationDtoToEducation(&s)
		educations = append(educations, *ed)
	}
	for i, s := range dto.Interests {
		fmt.Println(i, s)
		ed := mapInterestDtoToInterest(&s)
		interests = append(interests, *ed)
	}
	for i, s := range dto.Skills {
		fmt.Println(i, s)
		ed := mapSkillDtoToSkill(&s)
		skills = append(skills, *ed)
	}
	for i, s := range dto.Experiences {
		fmt.Println(i, s)
		ed := mapExperienceDtoToExperience(&s)

		experiences = append(experiences, *ed)
	}
	user.Skills = skills
	user.Educations = educations
	user.Experiences = experiences
	user.Interests = interests
	return user
}

func mapEducationDtoToEducation(e *dto.EducationDto) *model.Education {
	education := &model.Education{
		Education: e.Education,
	}
	return education
}
func mapSkillDtoToSkill(e *dto.SkillDto) *model.Skill {
	skill := &model.Skill{
		Skill: e.Skill,
	}
	return skill
}
func mapInterestDtoToInterest(e *dto.InterestDto) *model.Interest {
	interest := &model.Interest{
		Interest: e.Interest,
	}
	return interest
}
func mapExperienceDtoToExperience(e *dto.ExperienceDto) *model.Experience {
	experience := &model.Experience{
		Experience: e.Experience,
	}
	return experience
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
		ProfileStatus: userPb.UserRequest.ProfileStatus,
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

func MapPbToNewPasswordRequestDto(request *pb.NewPasswordRequest) *dto.NewRecoveryPasswordRequest {
	requestDto := &dto.NewRecoveryPasswordRequest{
		Username:    request.Recovery.Username,
		NewPassword: request.Recovery.Password,
		Code:        request.Recovery.Code,
	}
	return requestDto
}

func MapUserToUserDetails(user *model.User) *pb.UserDetails {
	var skills []*pb.Skill
	var interests []*pb.Interest
	var educations []*pb.Education
	var experiences []*pb.Experience

	for i, s := range user.Educations {
		fmt.Println(i, s)
		ed := mapEducationToEducationPb(&s)
		educations = append(educations, ed)
	}
	for i, s := range user.Interests {
		fmt.Println(i, s)
		ed := mapInterestToInterestPb(&s)
		interests = append(interests, ed)
	}
	for i, s := range user.Skills {
		fmt.Println(i, s)
		ed := mapSkillToSkillPb(&s)
		skills = append(skills, ed)
	}
	for i, s := range user.Experiences {
		fmt.Println(i, s)
		ed := mapExperienceToExperiencePb(&s)

		experiences = append(experiences, ed)
	}

	userDetails := &pb.UserDetails{
		Username:      user.Username,
		Email:         user.Email,
		PhoneNumber:   user.PhoneNumber,
		FirstName:     user.FirstName,
		LastName:      user.LastName,
		Gender:        mapGenderToString(user.Gender),
		DateOfBirth:   user.DateOfBirth.String(),
		Biography:     user.Biography,
		Skills:        skills,
		Interests:     interests,
		Experiences:   experiences,
		Educations:    educations,
		ProfileStatus: string(user.ProfileStatus),
	}
	return userDetails
}

func mapGenderToString(gender model.Gender) string {
	if gender == model.MALE {
		return "MALE"
	}
	if gender == model.FEMALE {
		return "FEMALE"
	} else {
		return "OTHER"
	}
}

func mapEducationToEducationPb(e *model.Education) *pb.Education {
	education := &pb.Education{
		Education: e.Education,
	}
	return education
}
func mapSkillToSkillPb(e *model.Skill) *pb.Skill {
	skill := &pb.Skill{
		Skill: e.Skill,
	}
	return skill
}
func mapInterestToInterestPb(e *model.Interest) *pb.Interest {
	interest := &pb.Interest{
		Interest: e.Interest,
	}
	return interest
}
func mapExperienceToExperiencePb(e *model.Experience) *pb.Experience {
	experience := &pb.Experience{
		Experience: e.Experience,
	}
	return experience
}

func MapUserPersonalDetailsDtoToUser(dto *dto.UserPersonalDetails, user *model.User) *model.User {
	user.Biography = dto.Biography
	user.FirstName = dto.FirstName
	user.LastName = dto.LastName
	user.Gender = mapGenderToModel(dto.Gender)
	user.PhoneNumber = dto.PhoneNumber
	user.DateOfBirth = mapToDate(dto.DateOfBirth)
	return user
}

func MapUserProfessionalDetailsDtoToUser(dto *dto.UserProfessionalDetails, user *model.User) *model.User {
	var skills []model.Skill
	var interests []model.Interest
	var educations []model.Education
	var experiences []model.Experience
	for i, s := range dto.Educations {
		fmt.Println(i, s)
		ed := mapEducationDtoToEducation(&s)
		educations = append(educations, *ed)
	}
	for i, s := range dto.Interests {
		fmt.Println(i, s)
		ed := mapInterestDtoToInterest(&s)
		interests = append(interests, *ed)
	}
	for i, s := range dto.Skills {
		fmt.Println(i, s)
		ed := mapSkillDtoToSkill(&s)
		skills = append(skills, *ed)
	}
	for i, s := range dto.Experiences {
		fmt.Println(i, s)
		ed := mapExperienceDtoToExperience(&s)

		experiences = append(experiences, *ed)
	}
	user.Skills = skills
	user.Educations = educations
	user.Experiences = experiences
	user.Interests = interests
	return user
}

func MapPbUserPersonalDetailsToUser(userPersonalDetailsPb *pb.UserPersonalDetailsRequest) *dto.UserPersonalDetails {

	userD := &dto.UserPersonalDetails{
		FirstName:   userPersonalDetailsPb.UserPersonalDetails.FirstName,
		LastName:    userPersonalDetailsPb.UserPersonalDetails.LastName,
		Username:    userPersonalDetailsPb.UserPersonalDetails.Username,
		PhoneNumber: userPersonalDetailsPb.UserPersonalDetails.PhoneNumber,
		Gender:      userPersonalDetailsPb.UserPersonalDetails.Gender,
		DateOfBirth: userPersonalDetailsPb.UserPersonalDetails.DateOfBirth,
		Biography:   userPersonalDetailsPb.UserPersonalDetails.Biography,
	}

	return userD
}

func MapUserToUserPersonalDetails(user *model.User) *pb.UserPersonalDetails {
	userPersonalDetails := &pb.UserPersonalDetails{
		Username:    user.Username,
		PhoneNumber: user.PhoneNumber,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Gender:      mapGenderToString(user.Gender),
		DateOfBirth: user.DateOfBirth.String(),
		Biography:   user.Biography,
	}
	return userPersonalDetails
}

func MapPbUserProfessionalDetailsToUser(userProfessionalDetailsPb *pb.UserProfessionalDetailsRequest) *dto.UserProfessionalDetails {

	userD := &dto.UserProfessionalDetails{
		Username: userProfessionalDetailsPb.UserProfessionalDetails.Username,
	}
	for i, s := range userProfessionalDetailsPb.UserProfessionalDetails.Educations {
		fmt.Println(i, s)
		ed := mapPbEducationToEducationDto(s)
		userD.Educations = append(userD.Educations, *ed)
	}
	for i, s := range userProfessionalDetailsPb.UserProfessionalDetails.Interests {
		fmt.Println(i, s)
		ed := mapPbEducationToInterestDto(s)
		userD.Interests = append(userD.Interests, *ed)
	}
	for i, s := range userProfessionalDetailsPb.UserProfessionalDetails.Skills {
		fmt.Println(i, s)
		ed := mapPbEducationToSkillDto(s)
		userD.Skills = append(userD.Skills, *ed)
	}
	for i, s := range userProfessionalDetailsPb.UserProfessionalDetails.Experiences {
		fmt.Println(i, s)
		ed := mapPbEducationToExperienceDto(s)
		userD.Experiences = append(userD.Experiences, *ed)
	}
	return userD
}

func MapUserToUserProfessionalDetails(user *model.User) *pb.UserProfessionalDetails {
	var skills []*pb.Skill
	var interests []*pb.Interest
	var educations []*pb.Education
	var experiences []*pb.Experience

	for i, s := range user.Educations {
		fmt.Println(i, s)
		ed := mapEducationToEducationPb(&s)
		educations = append(educations, ed)
	}
	for i, s := range user.Interests {
		fmt.Println(i, s)
		ed := mapInterestToInterestPb(&s)
		interests = append(interests, ed)
	}
	for i, s := range user.Skills {
		fmt.Println(i, s)
		ed := mapSkillToSkillPb(&s)
		skills = append(skills, ed)
	}
	for i, s := range user.Experiences {
		fmt.Println(i, s)
		ed := mapExperienceToExperiencePb(&s)

		experiences = append(experiences, ed)
	}

	userProfessionalDetails := &pb.UserProfessionalDetails{
		Username:    user.Username,
		Skills:      skills,
		Interests:   interests,
		Educations:  educations,
		Experiences: experiences,
	}
	return userProfessionalDetails
}

func MapToStringArrayInterests(interests []model.Interest) []string {
	var strings []string
	if len(interests) > 0 {
		for _, s := range interests {
			strings = append(strings, s.Interest)
		}
	}
	return strings
}

func MapToStringArraySkills(skills []model.Skill) []string {
	var strings []string
	if len(skills) > 0 {
		for _, s := range skills {
			strings = append(strings, s.Skill)
		}
	}
	return strings
}

func MapToStringArrayEducations(educations []model.Education) []string {
	var strings []string
	if len(educations) > 0 {
		for _, s := range educations {
			strings = append(strings, s.Education)
		}
	}
	return strings
}

func MapToStringArrayExperiences(experiences []model.Experience) []string {
	var strings []string
	if len(experiences) > 0 {
		for _, s := range experiences {
			strings = append(strings, s.Experience)
		}
	}
	return strings
}

func MapUserToEmailUsernameResponse(user *model.User) *pb.EmailUsernameResponse {
	emailUsername := &pb.EmailUsername{
		Email:    user.Email,
		Username: user.Username,
	}

	emailUsernameResponse := &pb.EmailUsernameResponse{
		UserId:        user.ID.String(),
		EmailUsername: emailUsername,
	}

	return emailUsernameResponse
}

func MapUserToChangeEmailResponse(user *model.User) *pb.ChangeEmailResponse {
	changeEmailResponse := &pb.ChangeEmailResponse{
		Email: user.Email,
	}
	return changeEmailResponse
}

func MapUserToChangeUsernameResponse(user *model.User) *pb.ChangeUsernameResponse {
	changeUsernameResponse := &pb.ChangeUsernameResponse{
		Username: user.Username,
	}
	return changeUsernameResponse
}
