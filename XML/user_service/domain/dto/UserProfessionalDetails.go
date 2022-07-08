package dto

type UserProfessionalDetails struct {
	Username    string          `json:"username"`
	Interests   []InterestDto   `json:"interests"`
	Skills      []SkillDto      `json:"skills"`
	Educations  []EducationDto  `json:"educations"`
	Experiences []ExperienceDto `json:"experiences"`
}
