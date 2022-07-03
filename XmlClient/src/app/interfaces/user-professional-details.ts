import { EducationDto } from "./education-dto";
import { ExperienceDto } from "./experience-dto";
import { InterestDto } from "./interest-dto";
import { SkillDto } from "./skill-dto";

export interface UserProfessionalDetails {
    username   : string;
	interests  : InterestDto[];
	skills     : SkillDto[];
	educations : EducationDto[];
	experiences : ExperienceDto[];
}
