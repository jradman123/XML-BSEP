import { Component, OnInit } from '@angular/core';
import { MatDialog, MatDialogConfig } from '@angular/material/dialog';
import { MatSnackBar } from '@angular/material/snack-bar';
import { Subscription } from 'rxjs';
import { EducationDto } from 'src/app/interfaces/education-dto';
import { ExperienceDto } from 'src/app/interfaces/experience-dto';
import { InterestDto } from 'src/app/interfaces/interest-dto';
import { SkillDto } from 'src/app/interfaces/skill-dto';
import { UserDetails } from 'src/app/interfaces/user-details';
import { UserService } from 'src/app/services/user-service/user.service';
import { EducationDialogComponent } from '../dialogs/education-dialog/education-dialog.component';

@Component({
  selector: 'app-profile-about',
  templateUrl: './profile-about.component.html',
  styleUrls: ['./profile-about.component.css']
})
export class ProfileAboutComponent implements OnInit {

  sub!: Subscription;
  userDetails! : UserDetails;
  id!: number;
  skills! : SkillDto[];
  educations! : EducationDto[];
  interests! : InterestDto[];
  experiences! : ExperienceDto[];
  newSkill! : SkillDto;
  newExperience! : ExperienceDto;
  newEducation! : EducationDto;
  newInterest! : InterestDto;
  skill! : string;
  interest! : string;
  education! : string;
  experience! : string;
  initialDetails : any;
  constructor(private userService : UserService,private matDialog : MatDialog,private _snackBar : MatSnackBar) {
      this.skills = [] as SkillDto[];
      this.educations = [] as EducationDto[];
      this.experiences = [] as ExperienceDto[];
      this.interests = [] as InterestDto[];
      this.newSkill = {} as SkillDto;
      this.newEducation = {} as EducationDto;
      this.newExperience = {} as ExperienceDto;
      this.newInterest = {} as InterestDto;
      this.education = "";
      this.skill = "";
      this.experience = "";
      this.interest = "";
   }

  ngOnInit(): void {
    this.getUserDetails();
  }
  getUserDetails() {
    this.sub = this.userService.getUserDetails(localStorage.getItem('username')).subscribe({
      next: (data: UserDetails) => {
        this.userDetails = data
        this.skills = data.skills;
        this.interests = data.interests;
        this.experiences = data.experiences;
        this.educations = data.educations;
      },
    });
  }

    openEducationDialog() {
      const dialogConfig = new MatDialogConfig();
      dialogConfig.disableClose = false;
      dialogConfig.id = 'modal-component';
      dialogConfig.height = 'fit-content';
      dialogConfig.width = '650px';
      const dialogRef = this.matDialog.open(EducationDialogComponent, dialogConfig);
      dialogRef.afterClosed().subscribe( data => {
        if (data) {
          this.education = data;
          this.save();
          }
        });
      
    }

    openExperienceDialog() {
      const dialogConfig = new MatDialogConfig();
      dialogConfig.disableClose = false;
      dialogConfig.id = 'modal-component';
      dialogConfig.height = 'fit-content';
      dialogConfig.width = '650px';
      const dialogRef = this.matDialog.open(EducationDialogComponent, dialogConfig);
      dialogRef.afterClosed().subscribe( data => {
        if (data) {
          this.experience = data;
          this.save();
          }
        });
      
    }

    openSkillDialog() {
      const dialogConfig = new MatDialogConfig();
      dialogConfig.disableClose = false;
      dialogConfig.id = 'modal-component';
      dialogConfig.height = 'fit-content';
      dialogConfig.width = '650px';
      const dialogRef = this.matDialog.open(EducationDialogComponent, dialogConfig);
      dialogRef.afterClosed().subscribe( data => {
        if (data) {
          this.skill = data;
          this.save();
          }
        });
      
    }

    openInterestDialog() {
      const dialogConfig = new MatDialogConfig();
      dialogConfig.disableClose = false;
      dialogConfig.id = 'modal-component';
      dialogConfig.height = 'fit-content';
      dialogConfig.width = '650px';
      const dialogRef = this.matDialog.open(EducationDialogComponent, dialogConfig);
      dialogRef.afterClosed().subscribe( data => {
        if (data) {
          this.interest = data;
          this.save();
          }
        });
      
    }

    save() {
      this.createUserDetails();
      
      this.userService.updateUser(this.userDetails).subscribe({
        next: () => {
          this.skills = [] as SkillDto[];
        this.educations = [] as EducationDto[];
        this.experiences = [] as ExperienceDto[];
        this.interests = [] as InterestDto[];
        this.newSkill = {} as SkillDto;
        this.newEducation = {} as EducationDto;
        this.newExperience = {} as ExperienceDto;
        this.newInterest = {} as InterestDto;
        this.education = "";
        this.skill = "";
        this.experience = "";
        this.interest = "";
          this.getUserDetails();
        }});
  
    }
  
    createUserDetails(): void {
      this.newSkill.skill = this.skill
      if(this.skill == ""){
        this.userDetails.skills = this.skills;
      }else{
        this.skills.push(this.newSkill);
        this.userDetails.skills = this.skills;
      }
      this.newEducation.education = this.education
      if(this.newEducation.education == ""){
        this.userDetails.educations =this.educations;
      }else{
        this.educations.push(this.newEducation);
        this.userDetails.educations =this.educations;
      }
      this.newExperience.experience = this.experience
      if(this.newExperience.experience == ""){
        this.userDetails.experiences = this.experiences;
      }else{
        this.experiences.push(this.newExperience);
        this.userDetails.experiences = this.experiences;
      }
      this.newInterest.interest = this.interest
      if(this.newInterest.interest == ""){
        this.userDetails.interests = this.interests;
      }else{
        this.interests.push(this.newInterest);
        this.userDetails.interests = this.interests;
      }
    }
  }
