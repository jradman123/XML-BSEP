import { Component, Input, OnInit } from '@angular/core';
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

    
  @Input()
  user! : UserDetails;
  sub!: Subscription;

  storageUsername : string | null = '';

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
    this.storageUsername = localStorage.getItem('username');
    this.setUserDetails();
  }
  
  setUserDetails() {
        this.skills = this.user.skills;
        this.interests = this.user.interests;
        this.experiences = this.user.experiences;
        this.educations = this.user.educations;
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
      
      this.userService.updateUser(this.user).subscribe({
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
          this.setUserDetails();
        }});
  
    }
  
    createUserDetails(): void {
      this.newSkill.skill = this.skill
      if(this.skill == ""){
        this.user.skills = this.skills;
      }else{
        this.skills.push(this.newSkill);
        this.user.skills = this.skills;
      }
      this.newEducation.education = this.education
      if(this.newEducation.education == ""){
        this.user.educations =this.educations;
      }else{
        this.educations.push(this.newEducation);
        this.user.educations =this.educations;
      }
      this.newExperience.experience = this.experience
      if(this.newExperience.experience == ""){
        this.user.experiences = this.experiences;
      }else{
        this.experiences.push(this.newExperience);
        this.user.experiences = this.experiences;
      }
      this.newInterest.interest = this.interest
      if(this.newInterest.interest == ""){
        this.user.interests = this.interests;
      }else{
        this.interests.push(this.newInterest);
        this.user.interests = this.interests;
      }
    }
  }
