package dto

type UserDetails struct {
	Username    string          `json:"username"`
	Email       string          `json:"email" validate:"required,email"`
	PhoneNumber string          `json:"phoneNumber" validate:"required"`
	FirstName   string          `json:"firstName" validate:"required,alpha,min=2,max=20"`
	LastName    string          `json:"lastName" validate:"required,alpha,min=2,max=35"`
	Gender      string          `json:"gender" validate:"oneof=MALE FEMALE OTHER"`
	DateOfBirth string          `json:"dateOfBirth" validate:"required"`
	Biography   string          `json:"biography"`
	Interests   []InterestDto   `json:"interests"`
	Skills      []SkillDto      `json:"skills"`
	Educations  []EducationDto  `json:"educations"`
	Experiences []ExperienceDto `json:"experiences"`
}
type SkillDto struct {
	Skill string `json:"skill"`
}
type InterestDto struct {
	Interest string `json:"interest"`
}
type EducationDto struct {
	Education string `json:"education"`
}
type ExperienceDto struct {
	Experience string `json:"experience"`
}
