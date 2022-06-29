import { EducationDto } from "./education-dto";
import { ExperienceDto } from "./experience-dto";
import { InterestDto } from "./interest-dto";
import { SkillDto } from "./skill-dto";

export interface UserDetails {
    username : string;
	email : string;
	phoneNumber :  string;          
	firstName : string;         
	lastName :  string;          
	gender :   string ;         
	dateOfBirth : string;          
	biography : string;          
	interests :  InterestDto[];   
	skills    :  SkillDto[]  ;    
	educations : EducationDto[];  
	experiences : ExperienceDto[]; 
}
