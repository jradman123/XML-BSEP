package api

import (
	pb "common/module/proto/user_service"
	"context"
	"fmt"
	"github.com/google/uuid"
	tracer "monitoring/module"
	"time"
	"user/module/domain/dto"
	"user/module/domain/model"
)

func MapProduct(user *model.User, ctx context.Context) *pb.User {
	span := tracer.StartSpanFromContext(ctx, "mapProduct")
	defer span.Finish()

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

func MapUserToPbResponseUser(user *model.User, ctx context.Context) *pb.RegisteredUser {
	span := tracer.StartSpanFromContext(ctx, "mapUserToPbResponseUser")
	defer span.Finish()

	usersPb := &pb.RegisteredUser{
		Username:  user.Username,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}
	return usersPb
}
func MapDtoToUser(userPb *dto.NewUser, ctx context.Context) *model.User {
	span := tracer.StartSpanFromContext(ctx, "mapDtoToUser")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
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
		Gender:        mapGenderToModel(userPb.Gender, ctx),
		DateOfBirth:   mapToDate(userPb.DateOfBirth, ctx),
		Password:      userPb.Password,
		Role:          model.Regular,
		IsConfirmed:   false,
		RecoveryEmail: userPb.RecoveryEmail,
		ProfileStatus: status,
	}
	return userD
}

func MapPbUserDetailsToUser(userDetailsPb *pb.UserDetailsRequest, ctx context.Context) *dto.UserDetails {
	span := tracer.StartSpanFromContext(ctx, "mapPbUserDetailsToUser")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
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
		ed := mapPbEducationToEducationDto(s, ctx)
		userD.Educations = append(userD.Educations, *ed)
	}
	for i, s := range userDetailsPb.UserDetails.Interests {
		fmt.Println(i, s)
		ed := mapPbEducationToInterestDto(s, ctx)
		userD.Interests = append(userD.Interests, *ed)
	}
	for i, s := range userDetailsPb.UserDetails.Skills {
		fmt.Println(i, s)
		ed := mapPbEducationToSkillDto(s, ctx)
		userD.Skills = append(userD.Skills, *ed)
	}
	for i, s := range userDetailsPb.UserDetails.Experiences {
		fmt.Println(i, s)
		ed := mapPbEducationToExperienceDto(s, ctx)
		userD.Experiences = append(userD.Experiences, *ed)
	}
	return userD
}
func mapPbEducationToEducationDto(e *pb.Education, ctx context.Context) *dto.EducationDto {
	span := tracer.StartSpanFromContext(ctx, "mapPbEducationToEducationDto")
	defer span.Finish()

	education := &dto.EducationDto{
		Education: e.Education,
	}
	return education
}
func mapPbEducationToSkillDto(e *pb.Skill, ctx context.Context) *dto.SkillDto {
	span := tracer.StartSpanFromContext(ctx, "mapPbEducationToSkillDto")
	defer span.Finish()

	skill := &dto.SkillDto{
		Skill: e.Skill,
	}
	return skill
}
func mapPbEducationToInterestDto(e *pb.Interest, ctx context.Context) *dto.InterestDto {
	span := tracer.StartSpanFromContext(ctx, "mapPbEducationToInterestDto")
	defer span.Finish()
	interest := &dto.InterestDto{
		Interest: e.Interest,
	}
	return interest
}
func mapPbEducationToExperienceDto(e *pb.Experience, ctx context.Context) *dto.ExperienceDto {
	span := tracer.StartSpanFromContext(ctx, "mapPbEducationToExperienceDto")
	defer span.Finish()

	experience := &dto.ExperienceDto{
		Experience: e.Experience,
	}
	return experience
}

func MapUserDetailsDtoToUser(dto *dto.UserDetails, user *model.User, ctx context.Context) *model.User {
	span := tracer.StartSpanFromContext(ctx, "mapUserDetailsDtoToUser")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	user.Biography = dto.Biography
	user.FirstName = dto.FirstName
	user.LastName = dto.LastName
	user.Gender = mapGenderToModel(dto.Gender, ctx)
	user.PhoneNumber = dto.PhoneNumber
	user.DateOfBirth = mapToDate(dto.DateOfBirth, ctx)
	//KAKO SAD DA SVE NIZOVE PONISTIM DA BUDU PRAZNI NA POCETKU
	var skills []model.Skill
	var interests []model.Interest
	var educations []model.Education
	var experiences []model.Experience
	for i, s := range dto.Educations {
		fmt.Println(i, s)
		ed := mapEducationDtoToEducation(&s, ctx)
		educations = append(educations, *ed)
	}
	for i, s := range dto.Interests {
		fmt.Println(i, s)
		ed := mapInterestDtoToInterest(&s, ctx)
		interests = append(interests, *ed)
	}
	for i, s := range dto.Skills {
		fmt.Println(i, s)
		ed := mapSkillDtoToSkill(&s, ctx)
		skills = append(skills, *ed)
	}
	for i, s := range dto.Experiences {
		fmt.Println(i, s)
		ed := mapExperienceDtoToExperience(&s, ctx)

		experiences = append(experiences, *ed)
	}
	user.Skills = skills
	user.Educations = educations
	user.Experiences = experiences
	user.Interests = interests
	return user
}

func mapEducationDtoToEducation(e *dto.EducationDto, ctx context.Context) *model.Education {
	span := tracer.StartSpanFromContext(ctx, "mapEducationDtoToEducation")
	defer span.Finish()

	education := &model.Education{
		Education: e.Education,
	}
	return education
}
func mapSkillDtoToSkill(e *dto.SkillDto, ctx context.Context) *model.Skill {
	span := tracer.StartSpanFromContext(ctx, "mapSkillDtoToSkill")
	defer span.Finish()

	skill := &model.Skill{
		Skill: e.Skill,
	}
	return skill
}
func mapInterestDtoToInterest(e *dto.InterestDto, ctx context.Context) *model.Interest {
	span := tracer.StartSpanFromContext(ctx, "mapInterestDtoToInterest")
	defer span.Finish()

	interest := &model.Interest{
		Interest: e.Interest,
	}
	return interest
}
func mapExperienceDtoToExperience(e *dto.ExperienceDto, ctx context.Context) *model.Experience {
	span := tracer.StartSpanFromContext(ctx, "mapExperienceDtoToExperience")
	defer span.Finish()

	experience := &model.Experience{
		Experience: e.Experience,
	}
	return experience
}

func MapPbUserToNewUserDto(userPb *pb.RegisterUserRequest, ctx context.Context) *dto.NewUser {
	span := tracer.StartSpanFromContext(ctx, "mapProduct")
	defer span.Finish()
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

func mapToDate(birth string, ctx context.Context) time.Time {
	span := tracer.StartSpanFromContext(ctx, "mapToDate")
	defer span.Finish()

	layout := "2006-01-02T15:04:05.000Z"
	dateOfBirth, _ := time.Parse(layout, birth)
	return dateOfBirth

}
func mapGenderToModel(gender string, ctx context.Context) model.Gender {
	span := tracer.StartSpanFromContext(ctx, "mapGenderToModel")
	defer span.Finish()

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

func MapPbToUserActivateRequest(request *pb.ActivationRequest, ctx context.Context) *dto.UserActivateRequest {
	span := tracer.StartSpanFromContext(ctx, "mapPbToUserActivateRequest")
	defer span.Finish()
	requestDTO := &dto.UserActivateRequest{
		Code:     request.Account.Code,
		Username: request.Account.Username,
	}
	return requestDTO
}

func MapPbToNewPasswordRequestDto(request *pb.NewPasswordRequest, ctx context.Context) *dto.NewRecoveryPasswordRequest {
	span := tracer.StartSpanFromContext(ctx, "mapPbToNewPasswordRequestDto")
	defer span.Finish()

	requestDto := &dto.NewRecoveryPasswordRequest{
		Username:    request.Recovery.Username,
		NewPassword: request.Recovery.Password,
		Code:        request.Recovery.Code,
	}
	return requestDto
}

func MapUserToUserDetails(user *model.User, ctx context.Context) *pb.UserDetails {
	span := tracer.StartSpanFromContext(ctx, "mapUserToUserDetails")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	var skills []*pb.Skill
	var interests []*pb.Interest
	var educations []*pb.Education
	var experiences []*pb.Experience

	for i, s := range user.Educations {
		fmt.Println(i, s)
		ed := mapEducationToEducationPb(&s, ctx)
		educations = append(educations, ed)
	}
	for i, s := range user.Interests {
		fmt.Println(i, s)
		ed := mapInterestToInterestPb(&s, ctx)
		interests = append(interests, ed)
	}
	for i, s := range user.Skills {
		fmt.Println(i, s)
		ed := mapSkillToSkillPb(&s, ctx)
		skills = append(skills, ed)
	}
	for i, s := range user.Experiences {
		fmt.Println(i, s)
		ed := mapExperienceToExperiencePb(&s, ctx)

		experiences = append(experiences, ed)
	}

	userDetails := &pb.UserDetails{
		Username:      user.Username,
		Email:         user.Email,
		PhoneNumber:   user.PhoneNumber,
		FirstName:     user.FirstName,
		LastName:      user.LastName,
		Gender:        mapGenderToString(user.Gender, ctx),
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

func mapGenderToString(gender model.Gender, ctx context.Context) string {
	span := tracer.StartSpanFromContext(ctx, "mapGenderToString")
	defer span.Finish()
	if gender == model.MALE {
		return "MALE"
	}
	if gender == model.FEMALE {
		return "FEMALE"
	} else {
		return "OTHER"
	}
}

func mapEducationToEducationPb(e *model.Education, ctx context.Context) *pb.Education {
	span := tracer.StartSpanFromContext(ctx, "mapEducationToEducationPb")
	defer span.Finish()

	education := &pb.Education{
		Education: e.Education,
	}
	return education
}
func mapSkillToSkillPb(e *model.Skill, ctx context.Context) *pb.Skill {
	span := tracer.StartSpanFromContext(ctx, "mapSkillToSkillPb")
	defer span.Finish()
	skill := &pb.Skill{
		Skill: e.Skill,
	}
	return skill
}
func mapInterestToInterestPb(e *model.Interest, ctx context.Context) *pb.Interest {
	span := tracer.StartSpanFromContext(ctx, "mapInterestToInterestPb")
	defer span.Finish()

	interest := &pb.Interest{
		Interest: e.Interest,
	}
	return interest
}
func mapExperienceToExperiencePb(e *model.Experience, ctx context.Context) *pb.Experience {
	span := tracer.StartSpanFromContext(ctx, "mapExperienceToExperiencePb")
	defer span.Finish()

	experience := &pb.Experience{
		Experience: e.Experience,
	}
	return experience
}

func MapUserPersonalDetailsDtoToUser(dto *dto.UserPersonalDetails, user *model.User, ctx context.Context) *model.User {
	span := tracer.StartSpanFromContext(ctx, "mapUserPersonalDetailsDtoToUser")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	user.Biography = dto.Biography
	user.FirstName = dto.FirstName
	user.LastName = dto.LastName
	user.Gender = mapGenderToModel(dto.Gender, ctx)
	user.PhoneNumber = dto.PhoneNumber
	user.DateOfBirth = mapToDate(dto.DateOfBirth, ctx)
	return user
}

func MapUserProfessionalDetailsDtoToUser(dto *dto.UserProfessionalDetails, user *model.User, ctx context.Context) *model.User {
	span := tracer.StartSpanFromContext(ctx, "mapUserProfessionalDetailsDtoToUser")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	var skills []model.Skill
	var interests []model.Interest
	var educations []model.Education
	var experiences []model.Experience
	for i, s := range dto.Educations {
		fmt.Println(i, s)
		ed := mapEducationDtoToEducation(&s, ctx)
		educations = append(educations, *ed)
	}
	for i, s := range dto.Interests {
		fmt.Println(i, s)
		ed := mapInterestDtoToInterest(&s, ctx)
		interests = append(interests, *ed)
	}
	for i, s := range dto.Skills {
		fmt.Println(i, s)
		ed := mapSkillDtoToSkill(&s, ctx)
		skills = append(skills, *ed)
	}
	for i, s := range dto.Experiences {
		fmt.Println(i, s)
		ed := mapExperienceDtoToExperience(&s, ctx)

		experiences = append(experiences, *ed)
	}
	user.Skills = skills
	user.Educations = educations
	user.Experiences = experiences
	user.Interests = interests
	return user
}

func MapPbUserPersonalDetailsToUser(userPersonalDetailsPb *pb.UserPersonalDetailsRequest, ctx context.Context) *dto.UserPersonalDetails {
	span := tracer.StartSpanFromContext(ctx, "mapPbUserPersonalDetailsToUser")
	defer span.Finish()

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

func MapUserToUserPersonalDetails(user *model.User, ctx context.Context) *pb.UserPersonalDetails {
	span := tracer.StartSpanFromContext(ctx, "mapUserToUserPersonalDetails")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	userPersonalDetails := &pb.UserPersonalDetails{
		Username:    user.Username,
		PhoneNumber: user.PhoneNumber,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Gender:      mapGenderToString(user.Gender, ctx),
		DateOfBirth: user.DateOfBirth.String(),
		Biography:   user.Biography,
	}
	return userPersonalDetails
}

func MapPbUserProfessionalDetailsToUser(userProfessionalDetailsPb *pb.UserProfessionalDetailsRequest, ctx context.Context) *dto.UserProfessionalDetails {
	span := tracer.StartSpanFromContext(ctx, "mapPbUserProfessionalDetailsToUser")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	userD := &dto.UserProfessionalDetails{
		Username: userProfessionalDetailsPb.UserProfessionalDetails.Username,
	}
	for i, s := range userProfessionalDetailsPb.UserProfessionalDetails.Educations {
		fmt.Println(i, s)
		ed := mapPbEducationToEducationDto(s, ctx)
		userD.Educations = append(userD.Educations, *ed)
	}
	for i, s := range userProfessionalDetailsPb.UserProfessionalDetails.Interests {
		fmt.Println(i, s)
		ed := mapPbEducationToInterestDto(s, ctx)
		userD.Interests = append(userD.Interests, *ed)
	}
	for i, s := range userProfessionalDetailsPb.UserProfessionalDetails.Skills {
		fmt.Println(i, s)
		ed := mapPbEducationToSkillDto(s, ctx)
		userD.Skills = append(userD.Skills, *ed)
	}
	for i, s := range userProfessionalDetailsPb.UserProfessionalDetails.Experiences {
		fmt.Println(i, s)
		ed := mapPbEducationToExperienceDto(s, ctx)
		userD.Experiences = append(userD.Experiences, *ed)
	}
	return userD
}

func MapUserToUserProfessionalDetails(user *model.User, ctx context.Context) *pb.UserProfessionalDetails {
	span := tracer.StartSpanFromContext(ctx, "mapUserToUserProfessionalDetails")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	var skills []*pb.Skill
	var interests []*pb.Interest
	var educations []*pb.Education
	var experiences []*pb.Experience

	for i, s := range user.Educations {
		fmt.Println(i, s)
		ed := mapEducationToEducationPb(&s, ctx)
		educations = append(educations, ed)
	}
	for i, s := range user.Interests {
		fmt.Println(i, s)
		ed := mapInterestToInterestPb(&s, ctx)
		interests = append(interests, ed)
	}
	for i, s := range user.Skills {
		fmt.Println(i, s)
		ed := mapSkillToSkillPb(&s, ctx)
		skills = append(skills, ed)
	}
	for i, s := range user.Experiences {
		fmt.Println(i, s)
		ed := mapExperienceToExperiencePb(&s, ctx)

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

func MapUserToEmailUsernameResponse(user *model.User, ctx context.Context) *pb.EmailUsernameResponse {
	span := tracer.StartSpanFromContext(ctx, "mapUserToEmailUsernameResponse")
	defer span.Finish()

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

func MapUserToChangeEmailResponse(user *model.User, ctx context.Context) *pb.ChangeEmailResponse {
	span := tracer.StartSpanFromContext(ctx, "mapUserToChangeEmailResponse")
	defer span.Finish()

	changeEmailResponse := &pb.ChangeEmailResponse{
		Email: user.Email,
	}
	return changeEmailResponse
}

func MapUserToChangeUsernameResponse(user *model.User, ctx context.Context) *pb.ChangeUsernameResponse {
	span := tracer.StartSpanFromContext(ctx, "mapUserToChangeUsernameResponse")
	defer span.Finish()

	changeUsernameResponse := &pb.ChangeUsernameResponse{
		Username: user.Username,
	}
	return changeUsernameResponse
}
