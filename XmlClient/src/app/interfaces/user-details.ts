import { EducationDto } from "./education-dto";
import { ExperienceDto } from "./experience-dto";
import { InterestDto } from "./interest-dto";
import { SkillDto } from "./skill-dto";

export interface UserDetails {
    username : string;
	email : string;
	phoneNumber :  string;          //`json:"phoneNumber" validate:"required"`
	firstName : string;          //`json:"firstName" validate:"required,alpha,min=2,max=20"`
	lastName :  string;          //`json:"lastName" validate:"required,alpha,min=2,max=35"`
	gender :   string ;         //`json:"gender" validate:"oneof=MALE FEMALE OTHER"`
	dateOfBirth : string;          //`json:"dateOfBirth" validate:"required"`
	biography : string;          //`json:"biography"`
	interests :  InterestDto[];   //`json:"interests"`
	skills    :  SkillDto[]  ;    //`json:"skills"`
	educations : EducationDto[];  //`json:"educations"`
	experiences : ExperienceDto[]; //`json:"experiences"`
}
